package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

const sourceURL = "https://raw.githubusercontent.com/reearth/reearth-cms/main/server/schemas/integration.yml"

func main() {
	if err := updateSchema(); err != nil {
		panic(err)
	}
}

func updateSchema() error {
	res, err := http.Get(sourceURL)
	if err != nil {
		return fmt.Errorf("failed to get schema: %w", err)
	}
	defer res.Body.Close()

	var schema any
	if err := yaml.NewDecoder(res.Body).Decode(&schema); err != nil {
		return fmt.Errorf("failed to decode schema: %w", err)
	}

	schema = removeXGoType(schema)

	if err := yaml.NewEncoder(os.Stdout).Encode(schema); err != nil {
		return fmt.Errorf("failed to encode schema: %w", err)
	}

	return nil
}

func removeXGoType(schema any) any {
	switch schema := schema.(type) {
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
