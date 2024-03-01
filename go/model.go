package cms

import (
	"time"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

const (
	AssetArchiveExtractionStatusDone = "done"
	tag                              = "cms"
)

type Asset struct {
	ID                      string `json:"id,omitempty"`
	ProjectID               string `json:"projectId,omitempty"`
	URL                     string `json:"url,omitempty"`
	ContentType             string `json:"contentType,omitempty"`
	ArchiveExtractionStatus string `json:"archiveExtractionStatus,omitempty"`
	File                    *File  `json:"file,omitempty"`
}

type AssetUpload struct {
	URL         string `json:"url"`
	Token       string `json:"token"`
	ContentType string `json:"contentType"`
}

type File struct {
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	Path        string `json:"path"`
	Children    []File `json:"children"`
}

func (f File) Paths() []string {
	return filePaths(f)
}

func filePaths(f File) (p []string) {
	if len(f.Children) == 0 {
		p = append(p, f.Path)
	}
	p = append(p, lo.FlatMap(f.Children, func(f File, _ int) []string {
		return filePaths(f)
	})...)
	return p
}

func (a *Asset) Clone() *Asset {
	if a == nil {
		return nil
	}
	return &Asset{
		ID:                      a.ID,
		ProjectID:               a.ProjectID,
		URL:                     a.URL,
		ArchiveExtractionStatus: a.ArchiveExtractionStatus,
	}
}

func (a *Asset) ToPublic() *PublicAsset {
	if a == nil {
		return nil
	}
	return &PublicAsset{
		Type:                    "asset",
		ID:                      a.ID,
		URL:                     a.URL,
		ContentType:             a.ContentType,
		ArchiveExtractionStatus: a.ArchiveExtractionStatus,
	}
}

type Models struct {
	Models     []Model `json:"models"`
	Page       int     `json:"page"`
	PerPage    int     `json:"perPage"`
	TotalCount int     `json:"totalCount"`
}

func (r Models) HasNext() bool {
	if r.PerPage == 0 {
		return false
	}
	return r.TotalCount > r.Page*r.PerPage
}

type Model struct {
	ID               string    `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Key              string    `json:"key,omitempty"`
	Public           bool      `json:"public,omitempty"`
	ProjectID        string    `json:"projectId,omitempty"`
	SchemaID         string    `json:"schemaId,omitempty"`
	MetadataSchemaID string    `json:"metadataSchemaId,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	LastModified     time.Time `json:"lastModified,omitempty"`
}

type Items struct {
	Items      []Item `json:"items"`
	Page       int    `json:"page"`
	PerPage    int    `json:"perPage"`
	TotalCount int    `json:"totalCount"`
}

func (r Items) HasNext() bool {
	if r.PerPage == 0 {
		return false
	}
	return r.TotalCount > r.Page*r.PerPage
}

type Item struct {
	ID              string   `json:"id"`
	ModelID         string   `json:"modelId"`
	Fields          []*Field `json:"fields"`
	MetadataFields  []*Field `json:"metadataFields,omitempty"`
	ReferencedItems []*Item  `json:"referencedItems,omitempty"`
	OriginalItemID  *string  `json:"originalItemId,omitempty"`
	MetadataItemID  *string  `json:"metadataItemId,omitempty"`
	IsMetadata      bool     `json:"isMetadata,omitempty"`
}

func (i *Item) Clone() *Item {
	if i == nil {
		return nil
	}
	return &Item{
		ID:      i.ID,
		ModelID: i.ModelID,
		Fields:  slices.Clone(i.Fields),
	}
}

func (i *Item) Field(id string) *Field {
	return i.FieldByGroup(id, "")
}

func (i *Item) MetadataField(id string) *Field {
	return i.MetadataFieldByGroup(id, "")
}

func (i *Item) FieldByGroup(id, group string) *Field {
	f, ok := lo.Find(i.Fields, func(f *Field) bool {
		return f.ID == id && f.Group == group
	})
	if ok {
		return f
	}
	return nil
}

func (i *Item) MetadataFieldByGroup(id, group string) *Field {
	f, ok := lo.Find(i.MetadataFields, func(f *Field) bool {
		return f.ID == id && f.Group == group
	})
	if ok {
		return f
	}
	return nil
}

func (i *Item) FieldByKey(key string) *Field {
	return i.FieldByKeyAndGroup(key, "")
}

func (i *Item) FieldByKeyAndGroup(key, group string) *Field {
	f, ok := lo.Find(i.Fields, func(f *Field) bool {
		return f.Key == key && f.Group == group
	})
	if ok {
		return f
	}
	return nil
}

func (i *Item) MetadataFieldByKey(key string) *Field {
	return i.MetadataFieldByKeyAndGroup(key, "")
}

func (i *Item) MetadataFieldByKeyAndGroup(key, group string) *Field {
	f, ok := lo.Find(i.MetadataFields, func(f *Field) bool {
		return f.Key == key && f.Group == group
	})
	if ok {
		return f
	}
	return nil
}

func (i *Item) Group(g string) *Item {
	fields := lo.Map(lo.Filter(i.Fields, func(f *Field, _ int) bool {
		return f.Group == g
	}), func(f *Field, _ int) *Field {
		g := f.Clone()
		g.Group = ""
		return g
	})

	metadataFields := lo.Map(lo.Filter(i.MetadataFields, func(f *Field, _ int) bool {
		return f.Group == g
	}), func(f *Field, _ int) *Field {
		g := f.Clone()
		g.Group = ""
		return g
	})

	return &Item{
		ID:             g,
		ModelID:        i.ModelID,
		Fields:         fields,
		MetadataFields: metadataFields,
	}
}

type Field struct {
	ID    string `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Group string `json:"group,omitempty"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func (f *Field) GetValue() *Value {
	if f == nil {
		return nil
	}
	return &Value{value: f.Value}
}

func (f *Field) Clone() *Field {
	if f == nil {
		return nil
	}
	return &Field{
		ID:    f.ID,
		Type:  f.Type,
		Value: f.Value,
		Key:   f.Key,
	}
}

type FieldChangeType string

const (
	FieldChangeTypeCreate FieldChangeType = "add"
	FieldChangeTypeUpdate FieldChangeType = "update"
	FieldChangeTypeDelete FieldChangeType = "delete"
)

type FieldChange struct {
	ID            string          `json:"id,omitempty"`
	Type          FieldChangeType `json:"type"`
	PreviousValue any             `json:"previousValue"`
	CurrentValue  any             `json:"currentValue"`
}

func (f *FieldChange) GetPreviousValue() *Value {
	if f == nil {
		return nil
	}
	return &Value{value: f.PreviousValue}
}

func (f *FieldChange) GetCurrentValue() *Value {
	if f == nil {
		return nil
	}
	return &Value{value: f.CurrentValue}
}

type Schema struct {
	ID        string        `json:"id"`
	Fields    []SchemaField `json:"fields"`
	ProjectID string        `json:"projectId"`
	Meta      *Schema       `json:"meta"`
}

func (d Schema) FieldIDByKey(k string) string {
	f, ok := lo.Find(d.Fields, func(f SchemaField) bool {
		return f.Key == k
	})
	if !ok {
		return ""
	}
	return f.ID
}

type SchemaField struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Key  string `json:"key"`
}
