package cms

import (
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

type Tag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func TagFrom(j any) *Tag {
	if j == nil {
		return nil
	}

	t := Tag{}
	if id, ok := j.(map[string]any)["id"].(string); ok {
		t.ID = id
	}
	if name, ok := j.(map[string]any)["name"].(string); ok {
		t.Name = name
	}
	if color, ok := j.(map[string]any)["color"].(string); ok {
		t.Color = color
	}

	return &t
}

type Value struct {
	value any
}

func NewValue(value any) *Value {
	return &Value{value: value}
}

func (v *Value) Interface() any {
	if v == nil {
		return nil
	}
	return v.value
}

func (v *Value) String() *string {
	return getValue[string](v)
}

func (v *Value) Int() *int {
	return getValue[int](v)
}

func (v *Value) Float() *float64 {
	return getValue[float64](v)
}

func (v *Value) Bool() *bool {
	return getValue[bool](v)
}

func (v *Value) Strings() []string {
	return getValues[string](v)
}

func (v *Value) Ints() []int {
	return getValues[int](v)
}

func (v *Value) Floats() []float64 {
	return getValues[float64](v)
}

func (v *Value) Bools() []bool {
	return getValues[bool](v)
}

func (v *Value) Tag() *Tag {
	if v == nil {
		return nil
	}
	return TagFrom(v.value)
}

func (v *Value) Tags() []Tag {
	values, ok := v.value.([]any)
	if !ok {
		values2, ok := v.value.([]map[string]any)
		if !ok {
			return nil
		}

		values = make([]any, len(values2))
		for i, value := range values2 {
			values[i] = value
		}
	}

	res := make([]Tag, len(values))
	for i, value := range values {
		if t := TagFrom(value); t != nil {
			res[i] = *t
		}
	}

	return res
}

func (f *Value) JSON(j any) error {
	if f == nil {
		return nil
	}

	s := f.String()
	if s == nil {
		return nil
	}

	err := json.Unmarshal([]byte(*s), &j)
	return err
}

func (f *Value) JSONs(j any) error {
	if f == nil {
		return nil
	}

	values := f.Strings()
	if values == nil {
		return nil
	}

	if res, ok := j.(*[]any); ok {
		*res = make([]any, len(values))
		j = *res
	}

	res, ok := j.([]any)
	if !ok {
		return nil
	}

	if len(values) != len(res) {
		return fmt.Errorf("length of values and j must be same")
	}

	for i, v := range values {
		if err := json.Unmarshal([]byte(v), &res[i]); err != nil {
			return fmt.Errorf("unmarshal json error at index %d: %w", i, err)
		}
	}

	return nil
}

func getValue[T any](v *Value) *T {
	if v == nil {
		return nil
	}

	if v, ok := v.value.(T); ok {
		return &v
	}

	return nil
}

func getValues[T any](v *Value) []T {
	if v == nil {
		return nil
	}

	if v, ok := v.value.([]T); ok {
		return v
	}

	if v, ok := v.value.([]any); ok {
		return lo.FilterMap(v, func(e any, _ int) (T, bool) {
			s, ok := e.(T)
			return s, ok
		})
	}

	return nil
}