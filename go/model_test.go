package cms

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestItem_Group(t *testing.T) {
	item := Item{
		ID:      "xxx",
		ModelID: "xxx",
		Fields: []*Field{
			{Key: "aaa", Value: "bbb"},
			{Key: "bbb", Value: []string{"ccc", "bbb"}},
			{Key: "ccc", Value: []string{"a", "b"}},
			{Key: "ddd", Value: map[string]any{"a": "b"}},
			{Key: "ggg", Value: []string{"1", "2"}},
			{Key: "aaa", Group: "1", Value: "123"},
		},
		MetadataFields: []*Field{
			{Key: "eee", Value: "xxx"},
		},
	}

	g := item.Group("1")
	assert.Equal(t, &Item{
		ID:      "1",
		ModelID: "xxx",
		Fields: []*Field{
			{Key: "aaa", Value: "123"},
		},
		MetadataFields: []*Field{},
	}, g)
}

func TestItem_Field(t *testing.T) {
	assert.Equal(t, &Field{
		ID: "bbb", Value: "ccc", Type: "string",
	}, (&Item{
		Fields: []*Field{
			{ID: "aaa", Value: "bbb", Type: "string"},
			{ID: "bbb", Value: "ccc", Type: "string"},
		},
	}).Field("bbb"))
	assert.Nil(t, (&Item{
		Fields: []*Field{
			{ID: "aaa", Key: "bbb", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}).Field("ccc"))
}

func TestItem_MetadataField(t *testing.T) {
	assert.Equal(t, &Field{
		ID: "bbb", Value: "ccc", Type: "string",
	}, (&Item{
		MetadataFields: []*Field{
			{ID: "aaa", Value: "bbb", Type: "string"},
			{ID: "bbb", Value: "ccc", Type: "string"},
		},
	}).MetadataField("bbb"))
	assert.Nil(t, (&Item{
		MetadataFields: []*Field{
			{ID: "aaa", Key: "bbb", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}).MetadataField("ccc"))
}

func TestItem_FieldByKey(t *testing.T) {
	assert.Equal(t, &Field{
		ID: "bbb", Key: "ccc", Type: "string",
	}, (&Item{
		Fields: []*Field{
			{ID: "aaa", Key: "bbb", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}).FieldByKey("ccc"))
	assert.Nil(t, (&Item{
		Fields: []*Field{
			{ID: "aaa", Key: "aaa", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}).FieldByKey("bbb"))
}

func TestItem_MetadataFieldByKey(t *testing.T) {
	assert.Equal(t, &Field{
		ID: "bbb", Key: "ccc", Type: "string",
	}, (&Item{
		MetadataFields: []*Field{
			{ID: "aaa", Key: "bbb", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}).MetadataFieldByKey("ccc"))
	assert.Nil(t, (&Item{
		MetadataFields: []*Field{
			{ID: "aaa", Key: "aaa", Type: "string"},
			{ID: "bbb", Key: "ccc", Type: "string"},
		},
	}).MetadataFieldByKey("bbb"))
}

func TestField_ValueString(t *testing.T) {
	assert.Equal(t, lo.ToPtr("ccc"), (&Field{
		Value: "ccc",
	}).GetValue().String())
	assert.Nil(t, (&Field{
		Value: 1,
	}).GetValue().String())
}

func TestField_ValueStrings(t *testing.T) {
	assert.Equal(t, []string{"ccc", "ddd"}, (&Field{
		Value: []string{"ccc", "ddd"},
	}).GetValue().Strings())
	assert.Equal(t, []string{"ccc", "ddd"}, (&Field{
		Value: []any{"ccc", "ddd", 1},
	}).GetValue().Strings())
	assert.Nil(t, (&Field{
		Value: "ccc",
	}).GetValue().Strings())
}

func TestField_ValueBool(t *testing.T) {
	assert.Equal(t, lo.ToPtr(true), (&Field{
		Value: true,
	}).GetValue().Bool())
	assert.Nil(t, (&Field{
		Value: 1,
	}).GetValue().Bool())
}

func TestField_ValueInt(t *testing.T) {
	assert.Equal(t, lo.ToPtr(100), (&Field{
		Value: 100,
	}).GetValue().Int())
	assert.Nil(t, (&Field{
		Value: "100",
	}).GetValue().Int())
}

func TestField_ValueFloat(t *testing.T) {
	assert.Equal(t, lo.ToPtr(100.1), (&Field{
		Value: 100.1,
	}).GetValue().Float())
	assert.Nil(t, (&Field{
		Value: 100,
	}).GetValue().Float())
	assert.Nil(t, (&Field{
		Value: "100.1",
	}).GetValue().Float())
}

func TestField_ValueTag(t *testing.T) {
	assert.Equal(t, &Tag{
		ID:    "xxx",
		Name:  "tag",
		Color: "red",
	}, (&Field{
		Value: map[string]any{
			"id":    "xxx",
			"name":  "tag",
			"color": "red",
		},
	}).GetValue().Tag())
	assert.Nil(t, (&Field{
		Value: 100,
	}).GetValue().Tag())
}

func TestField_ValueTags(t *testing.T) {
	assert.Equal(t, []Tag{
		{ID: "xxx"},
		{ID: "yyy"},
	}, (&Field{
		Value: []any{
			map[string]any{"id": "xxx"}, map[string]any{"id": "yyy"},
		},
	}).GetValue().Tags())
	assert.Equal(t, []Tag{
		{ID: "xxx"},
		{ID: "yyy"},
	}, (&Field{
		Value: []map[string]any{
			{"id": "xxx"}, {"id": "yyy"},
		},
	}).GetValue().Tags())
	assert.Nil(t, (&Field{
		Value: map[string]any{
			"id":    "xxx",
			"name":  "tag",
			"color": "red",
		},
	}).GetValue().Tags())
}

func TestField_ValueJSON(t *testing.T) {
	var r any
	err := (&Field{
		Value: `{"foo":"bar"}`,
	}).GetValue().JSON(&r)
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"foo": "bar"}, r)
}

func TestField_ValueJSONs(t *testing.T) {
	r := []any{}
	err := (&Field{
		Value: []string{`{"foo":"bar"}`, `{"foo":"hoge"}`},
	}).GetValue().JSONs(&r)
	assert.NoError(t, err)
	assert.Equal(t, []any{map[string]any{"foo": "bar"}, map[string]any{"foo": "hoge"}}, r)

	r = []any{nil, nil}
	err = (&Field{
		Value: []string{`{"foo":"bar"}`, `{"foo":"hoge"}`},
	}).GetValue().JSONs(r)
	assert.NoError(t, err)
	assert.Equal(t, []any{map[string]any{"foo": "bar"}, map[string]any{"foo": "hoge"}}, r)
}

func TestItems_HasNext(t *testing.T) {
	assert.True(t, Items{Page: 1, PerPage: 50, TotalCount: 100}.HasNext())
	assert.False(t, Items{Page: 2, PerPage: 50, TotalCount: 100}.HasNext())
	assert.True(t, Items{Page: 1, PerPage: 10, TotalCount: 11}.HasNext())
	assert.False(t, Items{Page: 2, PerPage: 10, TotalCount: 11}.HasNext())
}

func TestFile_Paths(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, File{
		Path: "_",
		Children: []File{
			{Path: "a"},
			{Path: "_", Children: []File{{Path: "b"}}},
			{Path: "c"},
		},
	}.Paths())
}
