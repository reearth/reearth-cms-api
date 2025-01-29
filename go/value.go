package cms

import (
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

type MarshalCMS interface {
	MarshalCMS() any
}

type Tag struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func TagFrom(j any) *Tag {
	if j == nil {
		return nil
	}

	m, ok := j.(map[string]any)
	if !ok {
		if m2, ok := j.(map[any]any); ok {
			m = make(map[string]any, len(m2))
			for k, v := range m2 {
				if k, ok := k.(string); ok {
					m[k] = v
				}
			}
		}
		return nil
	}

	t := Tag{}
	if id, ok := m["id"].(string); ok {
		t.ID = id
	}
	if name, ok := m["name"].(string); ok {
		t.Name = name
	}
	if color, ok := m["color"].(string); ok {
		t.Color = color
	}

	return &t
}

func TagsFrom(j any) []Tag {
	s, ok := j.([]any)
	if !ok {
		s2, ok := j.([]map[string]any)
		if !ok {
			return nil
		}
		s = make([]any, len(s2))
		for i, e := range s2 {
			s[i] = e
		}
	}

	res := make([]Tag, len(s))
	for i, e := range s {
		if t := TagFrom(e); t != nil {
			res[i] = *t
		}
	}

	return res
}

func (t *Tag) MarshalCMS() any {
	if t == nil || t.ID == "" && t.Name == "" {
		return nil
	}
	if t.ID == "" {
		return t.Name
	}
	return t.ID
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

func (v *Value) Asset() *PublicAsset {
	if v == nil {
		return nil
	}
	return PublicAssetFrom(v.value)
}

func (v *Value) AssetID() string {
	a := v.Asset()
	if a == nil {
		return ""
	}
	return a.ID
}

func (v *Value) AssetURL() string {
	a := v.Asset()
	if a == nil {
		return ""
	}
	return a.URL
}

func (v *Value) Tag() *Tag {
	if v == nil {
		return nil
	}
	return TagFrom(v.value)
}

func (v *Value) Tags() []Tag {
	if v == nil {
		return nil
	}
	return TagsFrom(v.value)
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

func (v *Value) MarshalCMS() any {
	if v == nil {
		return nil
	}
	if m, ok := v.value.(MarshalCMS); ok {
		return m.MarshalCMS()
	}
	return v.value
}

func (v *Value) MarshalJSON() ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(v.value)
}

func (v *Value) UnmarshalJSON(b []byte) error {
	if v == nil {
		return nil
	}

	if err := json.Unmarshal(b, &v.value); err != nil {
		return err
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

func PublicAssetFrom(a any) *PublicAsset {
	j, err := json.Marshal(a)
	if err != nil {
		return nil
	}

	pa := PublicAsset{}
	if err := json.Unmarshal(j, &pa); err != nil {
		return nil
	}

	return &pa
}

func PublicAssetsFrom(j any) []PublicAsset {
	s, ok := j.([]any)
	if !ok {
		s2, ok := j.([]map[string]any)
		if !ok {
			return nil
		}
		s = make([]any, len(s2))
		for i, e := range s2 {
			s[i] = e
		}
	}

	res := make([]PublicAsset, len(s))
	for i, e := range s {
		if t := PublicAssetFrom(e); t != nil {
			res[i] = *t
		}
	}

	return res
}

func AssetFrom(a any) *Asset {
	j, err := json.Marshal(a)
	if err != nil {
		return nil
	}

	pa := Asset{}
	if err := json.Unmarshal(j, &pa); err != nil {
		return nil
	}

	return &pa
}

func AssetsFrom(j any) []Asset {
	s, ok := j.([]any)
	if !ok {
		s2, ok := j.([]map[string]any)
		if !ok {
			return nil
		}
		s = make([]any, len(s2))
		for i, e := range s2 {
			s[i] = e
		}
	}

	res := make([]Asset, len(s))
	for i, e := range s {
		if t := AssetFrom(e); t != nil {
			res[i] = *t
		}
	}

	return res
}
