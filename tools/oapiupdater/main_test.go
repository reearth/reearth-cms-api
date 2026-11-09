package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveXGoType(t *testing.T) {
	input := map[string]any{
		"foo": "bar",
		"baz": map[string]any{
			"qux": "quux",
		},
		"quux": []any{
			"corge",
			"grault",
			map[string]any{
				"x-go-type": "garply",
			},
		},
	}

	expected := map[string]any{
		"foo": "bar",
		"baz": map[string]any{
			"qux": "quux",
		},
		"quux": []any{
			"corge",
			"grault",
			map[string]any{},
		},
	}

	assert.Equal(t, expected, removeXGoType(input))
}

func TestFetchSchema(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			_, _ = w.Write([]byte("openapi: 3.0.0\npaths:\n  /a:\n    x-go-type: foo\n    get: {}\n"))
		case "/notschema":
			_, _ = w.Write([]byte("hello"))
		default:
			http.Error(w, "404: Not Found", http.StatusNotFound)
		}
	}))
	defer srv.Close()

	got, err := fetchSchema(srv.URL + "/ok")
	assert.NoError(t, err)
	assert.NotContains(t, string(got), "x-go-type")
	assert.Contains(t, string(got), "openapi: 3.0.0")

	_, err = fetchSchema(srv.URL + "/missing")
	assert.ErrorContains(t, err, "404")

	_, err = fetchSchema(srv.URL + "/notschema")
	assert.Error(t, err)
}

func TestWriteFileAtomic(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "out.yml")
	assert.NoError(t, os.WriteFile(p, []byte("old"), 0o644))

	assert.NoError(t, writeFileAtomic(p, []byte("new")))
	b, _ := os.ReadFile(p)
	assert.Equal(t, "new", string(b))

	entries, _ := os.ReadDir(dir)
	assert.Len(t, entries, 1)
}

func TestSourceURL(t *testing.T) {
	assert.Equal(t, "https://raw.githubusercontent.com/reearth/reearth-cms/abc123/server/schemas/integration/integration.yml", sourceURL("abc123"))
}
