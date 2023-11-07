package main

import (
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
