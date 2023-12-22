package cms

import (
	"encoding/json"

	"github.com/samber/lo"
)

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

func (f *Value) JSON() (any, error) {
	if f == nil {
		return nil, nil
	}
	s := f.String()
	if s == nil {
		return nil, nil
	}

	var j any
	err := json.Unmarshal([]byte(*s), &j)
	return j, err
}

func (v *Value) JSONs() ([]any, error) {
	if v == nil {
		return nil, nil
	}
	values := v.Strings()
	if values == nil {
		return nil, nil
	}

	var res []any
	for _, v := range values {
		var r any
		if err := json.Unmarshal([]byte(v), &r); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
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
