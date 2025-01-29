package cms

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItem_Unmarshal(t *testing.T) {
	type str string

	type G struct {
		ID  string `cms:"id"`
		AAA string `cms:"aaa,text"`
	}

	type S struct {
		ID  string         `cms:"id"`
		AAA str            `cms:"aaa,"`
		BBB []string       `cms:"bbb"`
		CCC []str          `cms:"ccc"`
		DDD map[string]any `cms:"ddd"`
		EEE bool           `cms:"eee,,metadata"`
		GGG []*G           `cms:"ggg,group"`
		HHH []G            `cms:"hhh,group"`
		III *int           `cms:"iii,,metadata,includezero"`
		JJJ []Tag          `cms:"jjj"`
		KKK *Value         `cms:"kkk"`
		LLL *Tag           `cms:"lll"`
		MMM []string       `cms:"mmm"`
		NNN *PublicAsset   `cms:"nnn"`
	}
	s := S{}

	(&Item{
		ID: "xxx",
		Fields: []*Field{
			{Key: "aaa", Value: str("bbb")},
			{Key: "bbb", Value: []string{"ccc", "bbb"}},
			{Key: "ccc", Value: []str{"a", "b"}},
			{Key: "ddd", Value: map[string]any{"a": "b"}},
			{Key: "ggg", Type: "group", Value: []string{"1", "2"}},
			{Key: "hhh", Type: "group", Value: []string{"1"}},
			{Key: "aaa", Group: "1", Value: "123"},
			{Key: "iii"},
			{Key: "jjj", Value: []any{map[string]any{"id": "xxx!", "name": "tag"}}},
			{Key: "kkk", Value: []any{map[string]any{"id": "xxx", "name": "tag"}}},
			{Key: "lll", Value: map[string]any{"id": "xxx", "name": "tag"}},
			{Key: "mmm", Value: []any{"aaa", "bbb"}},
			{Key: "nnn", Value: map[string]any{"id": "asset", "url": "https://example.com"}},
		},
		MetadataFields: []*Field{
			{Key: "eee", Value: true},
		},
	}).Unmarshal(&s)

	assert.Equal(t, S{
		ID:  "xxx",
		AAA: str("bbb"),
		BBB: []string{"ccc", "bbb"},
		CCC: []str{"a", "b"},
		DDD: map[string]any{"a": "b"},
		EEE: true,
		GGG: []*G{{ID: "1", AAA: "123"}, {ID: "2"}},
		HHH: []G{{ID: "1", AAA: "123"}},
		III: nil,
		JJJ: []Tag{{ID: "xxx!", Name: "tag"}},
		KKK: &Value{value: []any{map[string]any{"id": "xxx", "name": "tag"}}},
		LLL: &Tag{ID: "xxx", Name: "tag"},
		MMM: []string{"aaa", "bbb"},
		NNN: &PublicAsset{Asset: Asset{ID: "asset", URL: "https://example.com"}},
	}, s)

	// no panic
	(&Item{}).Unmarshal(nil)
	(&Item{}).Unmarshal((*S)(nil))
}

func TestMarshal(t *testing.T) {
	type str string

	type G struct {
		ID  string `cms:"id"`
		AAA string `cms:"aaa,text"`
	}

	type S struct {
		ID  string   `cms:"id"`
		AAA string   `cms:"aaa,text"`
		BBB []string `cms:"bbb,select"`
		CCC str      `cms:"ccc"`
		DDD []str    `cms:"ddd"`
		EEE string   `cms:"eee,text"`
		FFF bool     `cms:"fff,bool,metadata"`
		GGG []G      `cms:"ggg"`
		HHH []*G     `cms:"hhh"`
		III *int     `cms:"iii,,metadata,includezero"`
		JJJ *Value   `cms:"jjj"`
		KKK []*Tag   `cms:"kkk"`
	}

	s := S{
		ID:  "xxx",
		AAA: "bbb",
		BBB: []string{"ccc", "bbb"},
		CCC: str("x"),
		DDD: []str{"1", "2"},
		FFF: true,
		GGG: []G{{ID: "1", AAA: "ggg"}},
		HHH: []*G{{ID: "2", AAA: "hhh"}, nil},
		JJJ: &Value{value: "foo"},
		KKK: []*Tag{{Name: "tag"}, nil},
	}

	expected := &Item{
		ID: "xxx",
		Fields: []*Field{
			{Key: "aaa", Type: "text", Value: "bbb"},
			{Key: "bbb", Type: "select", Value: []string{"ccc", "bbb"}},
			{Key: "ccc", Type: "", Value: str("x")},
			{Key: "ddd", Type: "", Value: []string{"1", "2"}},
			// no field for eee
			{Key: "aaa", Group: "1", Type: "text", Value: "ggg"},
			{Key: "ggg", Type: "group", Value: []string{"1"}},
			{Key: "aaa", Group: "2", Type: "text", Value: "hhh"},
			{Key: "hhh", Type: "group", Value: []string{"2"}},
			{Key: "jjj", Type: "", Value: "foo"},
			{Key: "kkk", Type: "", Value: []any{"tag"}},
		},
		MetadataFields: []*Field{
			{Key: "fff", Type: "bool", Value: true},
			{Key: "iii", Type: "", Value: (*int)(nil)},
		},
	}

	item := &Item{}
	Marshal(s, item)
	assert.Equal(t, expected, item)

	item2 := &Item{}
	Marshal(&s, item2)
	assert.Equal(t, item, item2)

	// no panic
	Marshal(nil, nil)
	Marshal(nil, item2)
	Marshal((*S)(nil), item2)
	Marshal(s, nil)
}

func TestConvertPrimitive(t *testing.T) {
	r, ok := convertPrimitive(reflect.ValueOf("a"), reflect.TypeOf(""))
	assert.True(t, ok)
	assert.Equal(t, reflect.ValueOf("a"), r)

	r, ok = convertPrimitive(reflect.ValueOf(any("a")), reflect.TypeOf(""))
	assert.True(t, ok)
	assert.Equal(t, reflect.ValueOf("a"), r)

	r, ok = convertPrimitive(reflect.ValueOf([]string{"a"}), reflect.TypeOf([]string{}))
	assert.True(t, ok)
	assert.Equal(t, []string{"a"}, r.Interface())

	r, ok = convertPrimitive(reflect.ValueOf([]any{"a"}), reflect.TypeOf([]string{}))
	assert.True(t, ok)
	assert.Equal(t, []string{"a"}, r.Interface())

	r, ok = convertPrimitive(reflect.ValueOf(map[string]any{"a": "b"}), reflect.TypeOf(map[string]string{}))
	assert.True(t, ok)
	assert.Equal(t, map[string]string{"a": "b"}, r.Interface())

	type s struct{ A string }
	r, ok = convertPrimitive(reflect.ValueOf(s{A: "a"}), reflect.TypeOf(s{}))
	assert.True(t, ok)
	assert.Equal(t, s{A: "a"}, r.Interface())

	r, ok = convertPrimitive(reflect.ValueOf(&s{A: "a"}), reflect.TypeOf(&s{}))
	assert.True(t, ok)
	assert.Equal(t, &s{A: "a"}, r.Interface())

	r, ok = convertPrimitive(reflect.ValueOf([]*s{{A: "a"}}), reflect.TypeOf([]*s{}))
	assert.True(t, ok)
	assert.Equal(t, []*s{{A: "a"}}, r.Interface())

	r, ok = convertPrimitive(reflect.ValueOf(true), reflect.TypeOf(""))
	assert.False(t, ok)
	assert.Zero(t, r)

	r, ok = convertPrimitive(reflect.ValueOf([]string{"a"}), reflect.TypeOf(""))
	assert.False(t, ok)
	assert.Zero(t, r)

	r, ok = convertPrimitive(reflect.ValueOf(any(nil)), reflect.TypeOf(""))
	assert.False(t, ok)
	assert.Zero(t, r)

	r, ok = convertPrimitive(reflect.ValueOf(map[string]string{"a": "a"}), reflect.TypeOf(map[bool]string{}))
	assert.False(t, ok)
	assert.Zero(t, r)
}
