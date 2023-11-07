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
	assert.Equal(t, Item{
		ID:      "1",
		ModelID: "xxx",
		Fields: []*Field{
			{Key: "aaa", Value: "123"},
		},
		MetadataFields: []*Field{},
	}, g)
}

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
	}
	s := S{}

	(&Item{
		ID: "xxx",
		Fields: []*Field{
			{Key: "aaa", Value: "bbb"},
			{Key: "bbb", Value: []string{"ccc", "bbb"}},
			{Key: "ccc", Value: []string{"a", "b"}},
			{Key: "ddd", Value: map[string]any{"a": "b"}},
			{Key: "ggg", Type: "group", Value: []string{"1", "2"}},
			{Key: "hhh", Type: "group", Value: []string{"1"}},
			{Key: "aaa", Group: "1", Value: "123"},
		},
		MetadataFields: []*Field{
			{Key: "eee", Value: true},
		},
	}).Unmarshal(&s)

	assert.Equal(t, S{
		ID:  "xxx",
		AAA: "bbb",
		BBB: []string{"ccc", "bbb"},
		CCC: []str{"a", "b"},
		DDD: map[string]any{"a": "b"},
		EEE: true,
		GGG: []*G{{ID: "1", AAA: "123"}, {ID: "2"}},
		HHH: []G{{ID: "1", AAA: "123"}},
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

func TestField_ValueJSON(t *testing.T) {
	r, err := (&Field{
		Value: `{"foo":"bar"}`,
	}).GetValue().JSON()
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"foo": "bar"}, r)
}

func TestField_ValueJSONs(t *testing.T) {
	r, err := (&Field{
		Value: []string{`{"foo":"bar"}`, `{"foo":"hoge"}`},
	}).GetValue().JSONs()
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
