package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	sourceURLTemplate = "https://raw.githubusercontent.com/reearth/reearth-cms/%s/server/schemas/integration/integration.yml"
	defaultRef        = "main"
)

func main() {
	out := flag.String("o", "", "output file (written atomically). If empty, writes to stdout")
	flag.Parse()

	ref := os.Getenv("REF")
	if ref == "" {
		ref = defaultRef
	}

	if err := run(ref, *out); err != nil {
		fmt.Fprintf(os.Stderr, "oapiupdater: %v\n", err)
		os.Exit(1)
	}
}

func sourceURL(ref string) string {
	return fmt.Sprintf(sourceURLTemplate, ref)
}

func run(ref, out string) error {
	url := sourceURL(ref)
	fmt.Fprintf(os.Stderr, "fetching %s\n", url)

	data, err := fetchSchema(url)
	if err != nil {
		return err
	}

	if out == "" {
		_, err := os.Stdout.Write(data)
		return err
	}
	return writeFileAtomic(out, data)
}

func fetchSchema(url string) ([]byte, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema: %w", err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get schema: unexpected status %s from %s", res.Status, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	return convertSchema(body)
}

func convertSchema(body []byte) ([]byte, error) {
	var schema any
	if err := yaml.Unmarshal(body, &schema); err != nil {
		return nil, fmt.Errorf("failed to decode schema: %w", err)
	}

	if m, ok := schema.(map[any]any); !ok || m["openapi"] == nil || m["paths"] == nil {
		return nil, fmt.Errorf("fetched document does not look like an OpenAPI schema")
	}

	schema = removeXGoType(schema)

	var buf bytes.Buffer
	if err := yaml.NewEncoder(&buf).Encode(schema); err != nil {
		return nil, fmt.Errorf("failed to encode schema: %w", err)
	}
	return buf.Bytes(), nil
}

// writeFileAtomic writes data to a temp file in the same directory and renames it
// over path, so path is never left empty or partially written on failure.
func writeFileAtomic(path string, data []byte) (err error) {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpName := tmp.Name()
	defer func() {
		if err != nil {
			_ = tmp.Close()
			_ = os.Remove(tmpName)
		}
	}()

	if _, err = tmp.Write(data); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}
	if err = tmp.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}
	if err = tmp.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	mode := os.FileMode(0o644)
	if fi, statErr := os.Stat(path); statErr == nil {
		mode = fi.Mode().Perm()
	}
	if err = os.Chmod(tmpName, mode); err != nil {
		return fmt.Errorf("failed to chmod temp file: %w", err)
	}

	if err = os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("failed to replace %s: %w", path, err)
	}
	return nil
}

func removeXGoType(schema any) any {
	switch schema := schema.(type) {
	case map[string]any:
		delete(schema, "x-go-type")
		for k, v := range schema {
			schema[k] = removeXGoType(v)
		}
	case map[any]any:
		delete(schema, "x-go-type")
		for k, v := range schema {
			schema[k] = removeXGoType(v)
		}
	case []any:
		for k, v := range schema {
			schema[k] = removeXGoType(v)
		}
	}
	return schema
}
