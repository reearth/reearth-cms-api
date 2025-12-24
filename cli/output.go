package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	cms "github.com/reearth/reearth-cms-api/go"
)

type OutputFormat int

const (
	OutputTable OutputFormat = iota
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
	return o.outputModelsTable(models)
}

func (o *Outputter) OutputModel(model *cms.Model) error {
	if o.format == OutputJSON {
		return o.outputJSON(model)
	}
	return o.outputModelDetail(model)
}

func (o *Outputter) outputModelDetail(model *cms.Model) error {
	_, _ = fmt.Fprintf(o.writer, "ID: %s\n", model.ID)
	_, _ = fmt.Fprintf(o.writer, "Name: %s\n", model.Name)
	_, _ = fmt.Fprintf(o.writer, "Key: %s\n", model.Key)
	_, _ = fmt.Fprintf(o.writer, "Public: %v\n", model.Public)
	_, _ = fmt.Fprintf(o.writer, "Project ID: %s\n", model.ProjectID)
	_, _ = fmt.Fprintf(o.writer, "Schema ID: %s\n", model.SchemaID)
	_, _ = fmt.Fprintf(o.writer, "Created At: %s\n", model.CreatedAt.Format("2006-01-02 15:04:05"))
	_, _ = fmt.Fprintf(o.writer, "Updated At: %s\n", model.UpdatedAt.Format("2006-01-02 15:04:05"))
	return nil
}

func (o *Outputter) outputModelsTable(models *cms.Models) error {
	table := tablewriter.NewWriter(o.writer)
	table.SetHeader([]string{"ID", "Name", "Key", "Public", "Created At"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, m := range models.Models {
		table.Append([]string{
			m.ID,
			m.Name,
			m.Key,
			fmt.Sprintf("%v", m.Public),
			m.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	table.Render()

	_, _ = fmt.Fprintf(o.writer, "\nPage %d (total: %d)\n", models.Page, models.TotalCount)
	return nil
}

func (o *Outputter) OutputItems(items *cms.Items) error {
	if o.format == OutputJSON {
		return o.outputJSON(items.Items)
	}
	return o.outputItemsTable(items)
}

func (o *Outputter) outputItemsTable(items *cms.Items) error {
	table := tablewriter.NewWriter(o.writer)
	table.SetHeader([]string{"ID", "Model ID", "Fields", "Created At"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, item := range items.Items {
		fieldCount := len(item.Fields)
		table.Append([]string{
			item.ID,
			item.ModelID,
			fmt.Sprintf("%d fields", fieldCount),
			item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	table.Render()

	_, _ = fmt.Fprintf(o.writer, "\nPage %d (total: %d)\n", items.Page, items.TotalCount)
	return nil
}

func (o *Outputter) OutputItem(item *cms.Item) error {
	if o.format == OutputJSON {
		return o.outputJSON(item)
	}
	return o.outputItemDetail(item)
}

func (o *Outputter) outputItemDetail(item *cms.Item) error {
	_, _ = fmt.Fprintf(o.writer, "ID: %s\n", item.ID)
	_, _ = fmt.Fprintf(o.writer, "Model ID: %s\n", item.ModelID)
	_, _ = fmt.Fprintf(o.writer, "Created At: %s\n", item.CreatedAt.Format("2006-01-02 15:04:05"))
	_, _ = fmt.Fprintf(o.writer, "Updated At: %s\n", item.UpdatedAt.Format("2006-01-02 15:04:05"))
	_, _ = fmt.Fprintf(o.writer, "\nFields:\n")

	table := tablewriter.NewWriter(o.writer)
	table.SetHeader([]string{"Key", "Type", "Value"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, f := range item.Fields {
		value := fmt.Sprintf("%v", f.Value)
		if len(value) > 50 {
			value = value[:50] + "..."
		}
		table.Append([]string{f.Key, f.Type, value})
	}
	table.Render()
	return nil
}

func (o *Outputter) OutputAsset(asset *cms.Asset) error {
	if o.format == OutputJSON {
		return o.outputJSON(asset)
	}
	return o.outputAssetDetail(asset)
}

func (o *Outputter) outputAssetDetail(asset *cms.Asset) error {
	_, _ = fmt.Fprintf(o.writer, "ID: %s\n", asset.ID)
	_, _ = fmt.Fprintf(o.writer, "Name: %s\n", asset.Name)
	_, _ = fmt.Fprintf(o.writer, "URL: %s\n", asset.URL)
	_, _ = fmt.Fprintf(o.writer, "Content Type: %s\n", asset.ContentType)
	_, _ = fmt.Fprintf(o.writer, "Project ID: %s\n", asset.ProjectID)
	_, _ = fmt.Fprintf(o.writer, "Created At: %s\n", asset.CreatedAt.Format("2006-01-02 15:04:05"))
	_, _ = fmt.Fprintf(o.writer, "Updated At: %s\n", asset.UpdatedAt.Format("2006-01-02 15:04:05"))
	return nil
}

func (o *Outputter) OutputMessage(msg string) {
	_, _ = fmt.Fprintln(o.writer, msg)
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
