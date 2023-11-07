package cms

import (
	"reflect"
	"strings"
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

type Model struct {
	ID           string    `json:"id"`
	Key          string    `json:"key,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
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

func (i *Item) Group(g string) Item {
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

	return Item{
		ID:             g,
		ModelID:        i.ModelID,
		Fields:         fields,
		MetadataFields: metadataFields,
	}
}

func (d *Item) Unmarshal(i any) {
	if i == nil {
		return
	}

	v := reflect.ValueOf(i)
	if v.IsNil() {
		return
	}

	v = v.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		key, opts, _ := strings.Cut(tag, ",")
		if key == "" || key == "-" {
			continue
		}

		isMetadata := strings.HasSuffix(opts, ",metadata")
		vf := v.FieldByName(f.Name)

		if key == "id" {
			if f.Type.Kind() == reflect.String {
				vf.SetString(d.ID)
			}
			continue
		}

		var itf *Field
		if isMetadata {
			itf = d.MetadataFieldByKey(key)
		} else {
			itf = d.FieldByKey(key)
		}

		if itf != nil && itf.Type == "group" {
			groupIDs := itf.GetValue().Strings()
			if len(groupIDs) == 0 {
				if groupID := itf.GetValue().String(); groupID != nil {
					groupIDs = []string{*groupID}
				}
			}

			groups := make([]Item, 0, len(groupIDs))
			for _, g := range groupIDs {
				group := d.Group(g)
				groups = append(groups, group)
			}

			if len(groups) == 0 {
				continue
			}

			if f.Type.Kind() == reflect.Slice &&
				(f.Type.Elem().Kind() == reflect.Struct ||
					f.Type.Elem().Kind() == reflect.Ptr &&
						f.Type.Elem().Elem().Kind() == reflect.Struct) {
				s := reflect.MakeSlice(f.Type, 0, len(groups))
				isPointer := f.Type.Elem().Kind() == reflect.Ptr

				for _, g := range groups {
					var rv reflect.Value
					if isPointer {
						rv = reflect.New(f.Type.Elem().Elem())
					} else {
						rv = reflect.New(f.Type.Elem())
					}

					i := rv.Interface()
					g.Unmarshal(i)

					if isPointer {
						s = reflect.Append(s, rv)
					} else {
						s = reflect.Append(s, rv.Elem())
					}
				}

				vf.Set(s)
			} else if f.Type.Kind() == reflect.Struct {
				groups[0].Unmarshal(vf.Addr().Interface())
			} else if f.Type.Kind() == reflect.Pointer && f.Type.Elem().Kind() == reflect.Struct {
				groups[0].Unmarshal(vf.Interface())
			}
		}

		if itf == nil || !vf.CanSet() || !reflect.TypeOf(itf.Value).AssignableTo(vf.Type()) {
			continue
		}

		vf.Set(reflect.ValueOf(itf.Value))
	}
}

func Marshal(i any, item *Item) {
	if item == nil || i == nil {
		return
	}

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Pointer {
		if v.IsNil() {
			return
		}
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	ni := Item{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		key, opts, _ := strings.Cut(tag, ",")
		if key == "" || key == "-" {
			continue
		}

		ty, optsRemaining, _ := strings.Cut(opts, ",")
		optsSplited := strings.Split(optsRemaining, ",")
		omitempty := !slices.Contains(optsSplited, "includezero")
		isMetadata := slices.Contains(optsSplited, "metadata")

		vf := v.FieldByName(f.Name)
		if key == "id" {
			ni.ID, _ = vf.Interface().(string)
			continue
		}

		vft := vf.Type()
		var value any
		if vft.Kind() == reflect.Slice && vft.Elem().Kind() == reflect.String && vf.Len() > 0 {
			st := reflect.TypeOf("")
			v := make([]string, 0, vf.Len())
			for i := 0; i < cap(v); i++ {
				vfs := vf.Index(i).Convert(st)
				v = append(v, vfs.String())
			}
			value = v
		} else if vft.Kind() == reflect.Slice && vf.Len() > 0 && (vft.Elem().Kind() == reflect.Struct ||
			vft.Elem().Kind() == reflect.Ptr && vft.Elem().Elem().Kind() == reflect.Struct) {
			isPointer := vft.Elem().Kind() == reflect.Ptr

			v := make([]string, 0, vf.Len())
			for i := 0; i < cap(v); i++ {
				var in any
				if isPointer {
					in = vf.Index(i).Interface()
				} else {
					in = vf.Index(i).Addr().Interface()
				}

				item := Item{}
				Marshal(in, &item)
				if item.ID == "" {
					continue
				}

				// assign group
				for i := range item.Fields {
					item.Fields[i].Group = item.ID
				}
				for i := range item.MetadataFields {
					item.Fields[i].Group = item.ID
				}

				// merge i to ni
				ni.Fields = append(ni.Fields, item.Fields...)
				ni.MetadataFields = append(ni.MetadataFields, item.MetadataFields...)

				v = append(v, item.ID)
			}

			if len(v) > 0 {
				value = v
				ty = "group"
			}
		} else if vft.Kind() == reflect.Struct || vft.Kind() == reflect.Ptr && vft.Elem().Kind() == reflect.Struct {
			isPointer := vft.Kind() == reflect.Ptr
			var v any
			if isPointer {
				v = vf.Interface()
			} else {
				v = vf.Addr().Interface()
			}

			item := Item{}
			Marshal(v, &item)
			if item.ID == "" {
				continue
			}

			// assign group
			for i := range item.Fields {
				item.Fields[i].Group = item.ID
			}
			for i := range item.MetadataFields {
				item.Fields[i].Group = item.ID
			}

			// merge i to ni
			ni.Fields = append(ni.Fields, item.Fields...)
			ni.MetadataFields = append(ni.MetadataFields, item.MetadataFields...)

			value = item.ID
			ty = "group"
		} else if !omitempty || !vf.IsZero() {
			value = vf.Convert(vft).Interface()
		}

		if value != nil {
			f := &Field{
				Key:   key,
				Type:  ty,
				Value: value,
			}

			if isMetadata {
				ni.MetadataFields = append(ni.MetadataFields, f)
			} else {
				ni.Fields = append(ni.Fields, f)
			}
		}
	}

	*item = ni
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
