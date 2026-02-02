package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	cms "github.com/reearth/reearth-cms-api/go"
	"gopkg.in/yaml.v3"
)

type OutputFormat int

const (
	OutputYAML OutputFormat = iota
	OutputJSON
)

type Outputter struct {
	format OutputFormat
	fields []string
	writer io.Writer
}

func NewOutputter(jsonFlag string) *Outputter {
	o := &Outputter{writer: os.Stdout}
	if jsonFlag != "" {
		o.format = OutputJSON
		if jsonFlag != "true" && jsonFlag != "1" {
			o.fields = strings.Split(jsonFlag, ",")
		}
	}
	return o
}

func (o *Outputter) OutputModels(models *cms.Models) error {
	if o.format == OutputJSON {
		return o.outputJSON(models.Models)
	}
	return o.outputModelsYAML(models)
}

func (o *Outputter) OutputModel(model *cms.Model) error {
	if o.format == OutputJSON {
		return o.outputJSON(model)
	}
	return o.outputModelYAML(model)
}

func (o *Outputter) outputModelYAML(model *cms.Model) error {
	data := modelToMap(model)
	return o.outputYAML(data)
}

func (o *Outputter) outputModelsYAML(models *cms.Models) error {
	items := make([]map[string]any, 0, len(models.Models))
	for _, m := range models.Models {
		items = append(items, modelToMap(&m))
	}

	if err := o.outputYAML(items); err != nil {
		return err
	}

	_, _ = fmt.Fprintf(o.writer, "\npage: %d\ntotal: %d\n", models.Page, models.TotalCount)
	return nil
}

func modelToMap(m *cms.Model) map[string]any {
	return map[string]any{
		"id":        m.ID,
		"name":      m.Name,
		"key":       m.Key,
		"public":    m.Public,
		"projectId": m.ProjectID,
		"schemaId":  m.SchemaID,
		"createdAt": formatTime(m.CreatedAt),
		"updatedAt": formatTime(m.UpdatedAt),
	}
}

func (o *Outputter) OutputItems(items *cms.Items) error {
	if o.format == OutputJSON {
		return o.outputJSON(items.Items)
	}
	return o.outputItemsYAML(items)
}

func (o *Outputter) outputItemsYAML(items *cms.Items) error {
	list := make([]map[string]any, 0, len(items.Items))
	for _, item := range items.Items {
		list = append(list, itemToMap(&item))
	}

	if err := o.outputYAML(list); err != nil {
		return err
	}

	_, _ = fmt.Fprintf(o.writer, "\npage: %d\ntotal: %d\n", items.Page, items.TotalCount)
	return nil
}

func (o *Outputter) OutputItem(item *cms.Item) error {
	if o.format == OutputJSON {
		return o.outputJSON(item)
	}
	return o.outputItemYAML(item)
}

func (o *Outputter) outputItemYAML(item *cms.Item) error {
	data := itemToMap(item)
	return o.outputYAML(data)
}

func itemToMap(item *cms.Item) map[string]any {
	// Group ID -> フィールドのマップを作成
	groupFields := make(map[string][]*cms.Field)
	for _, f := range item.Fields {
		if f.Group != "" {
			groupFields[f.Group] = append(groupFields[f.Group], f)
		}
	}

	// トップレベルフィールドを処理
	var topLevelFields []*cms.Field
	for _, f := range item.Fields {
		if f.Group == "" {
			topLevelFields = append(topLevelFields, f)
		}
	}

	return map[string]any{
		"id":        item.ID,
		"modelId":   item.ModelID,
		"createdAt": formatTime(item.CreatedAt),
		"updatedAt": formatTime(item.UpdatedAt),
		"fields":    fieldsToSlice(topLevelFields, groupFields),
	}
}

func fieldsToSlice(fields []*cms.Field, groupFields map[string][]*cms.Field) []map[string]any {
	result := make([]map[string]any, 0, len(fields))
	for _, f := range fields {
		if f.Type == "group" {
			// groupフィールドのvalueをネストしたフィールドに置き換え
			groupIDs := toStringSlice(f.Value)
			groupValue := make([][]map[string]any, 0, len(groupIDs))
			for _, gid := range groupIDs {
				if gf, ok := groupFields[gid]; ok {
					// 再帰的にネストしたグループを処理
					groupValue = append(groupValue, fieldsToSlice(gf, groupFields))
				}
			}
			result = append(result, map[string]any{
				"key":   f.Key,
				"type":  f.Type,
				"value": groupValue,
			})
		} else {
			result = append(result, map[string]any{
				"key":   f.Key,
				"type":  f.Type,
				"value": f.Value,
			})
		}
	}
	return result
}

func toStringSlice(v any) []string {
	if v == nil {
		return nil
	}
	// 単一の文字列の場合
	if s, ok := v.(string); ok {
		return []string{s}
	}
	if arr, ok := v.([]any); ok {
		result := make([]string, 0, len(arr))
		for _, item := range arr {
			if s, ok := item.(string); ok {
				result = append(result, s)
			}
		}
		return result
	}
	if arr, ok := v.([]string); ok {
		return arr
	}
	return nil
}

func (o *Outputter) OutputAsset(asset *cms.Asset) error {
	if o.format == OutputJSON {
		return o.outputJSON(asset)
	}
	return o.outputAssetYAML(asset)
}

func (o *Outputter) outputAssetYAML(asset *cms.Asset) error {
	data := map[string]any{
		"id":          asset.ID,
		"name":        asset.Name,
		"url":         asset.URL,
		"contentType": asset.ContentType,
		"projectId":   asset.ProjectID,
		"createdAt":   formatTime(asset.CreatedAt),
		"updatedAt":   formatTime(asset.UpdatedAt),
	}
	return o.outputYAML(data)
}

func (o *Outputter) OutputMessage(msg string) {
	_, _ = fmt.Fprintln(o.writer, msg)
}

func (o *Outputter) outputYAML(data any) error {
	enc := yaml.NewEncoder(o.writer)
	enc.SetIndent(2)
	defer enc.Close()
	return enc.Encode(data)
}

func (o *Outputter) outputJSON(data any) error {
	if len(o.fields) > 0 {
		data = filterFields(data, o.fields)
	}

	enc := json.NewEncoder(o.writer)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func filterFields(data any, fields []string) any {
	b, err := json.Marshal(data)
	if err != nil {
		return data
	}

	switch data.(type) {
	case []cms.Model, []cms.Item:
		var arr []map[string]any
		if err := json.Unmarshal(b, &arr); err != nil {
			return data
		}
		result := make([]map[string]any, len(arr))
		for i, item := range arr {
			result[i] = filterMap(item, fields)
		}
		return result
	default:
		var m map[string]any
		if err := json.Unmarshal(b, &m); err != nil {
			return data
		}
		return filterMap(m, fields)
	}
}

func filterMap(m map[string]any, fields []string) map[string]any {
	result := make(map[string]any)
	for _, f := range fields {
		f = strings.TrimSpace(f)
		if val, ok := m[f]; ok {
			result[f] = val
		}
	}
	return result
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}
