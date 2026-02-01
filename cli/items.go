package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	cms "github.com/reearth/reearth-cms-api/go"
	"github.com/spf13/cobra"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Manage Re:Earth CMS items",
}

var itemsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in a model",
	RunE:  runItemsList,
}

var itemsGetCmd = &cobra.Command{
	Use:   "get <item-id>",
	Short: "Get a specific item by ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runItemsGet,
}

var itemsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new item",
	RunE:  runItemsCreate,
}

var itemsUpdateCmd = &cobra.Command{
	Use:   "update <item-id>",
	Short: "Update an existing item",
	Args:  cobra.ExactArgs(1),
	RunE:  runItemsUpdate,
}

var itemsDeleteCmd = &cobra.Command{
	Use:   "delete <item-id>",
	Short: "Delete an item",
	Args:  cobra.ExactArgs(1),
	RunE:  runItemsDelete,
}

var (
	itemsModelID     string
	itemsPage        int
	itemsPerPage     int
	itemsAsset       bool
	itemsFieldIDs    []string
	itemsFieldKeys   []string
	itemsFieldTypes  []string
	itemsFieldValues []string
	itemsMetaIDs     []string
	itemsMetaKeys    []string
	itemsMetaTypes   []string
	itemsMetaValues  []string
	itemsYes         bool
)

func init() {
	itemsCmd.AddCommand(itemsListCmd)
	itemsCmd.AddCommand(itemsGetCmd)
	itemsCmd.AddCommand(itemsCreateCmd)
	itemsCmd.AddCommand(itemsUpdateCmd)
	itemsCmd.AddCommand(itemsDeleteCmd)

	// List flags
	itemsListCmd.Flags().StringVarP(&itemsModelID, "model", "m", "",
		"Model ID or key (required)")
	itemsListCmd.Flags().IntVar(&itemsPage, "page", 1, "Page number")
	itemsListCmd.Flags().IntVar(&itemsPerPage, "per-page", 50, "Items per page")
	itemsListCmd.Flags().BoolVar(&itemsAsset, "asset", false, "Include asset data")
	_ = itemsListCmd.MarkFlagRequired("model")

	// Get flags
	itemsGetCmd.Flags().BoolVar(&itemsAsset, "asset", false, "Include asset data")

	// Create flags
	itemsCreateCmd.Flags().StringVarP(&itemsModelID, "model", "m", "",
		"Model ID or key (required)")
	itemsCreateCmd.Flags().StringArrayVar(&itemsFieldIDs, "id", nil,
		"Field ID (optional, use with -k, -t, -v)")
	itemsCreateCmd.Flags().StringArrayVarP(&itemsFieldKeys, "key", "k", nil,
		"Field key (required for each field)")
	itemsCreateCmd.Flags().StringArrayVarP(&itemsFieldTypes, "type", "t", nil,
		"Field type (required for each field)")
	itemsCreateCmd.Flags().StringArrayVarP(&itemsFieldValues, "value", "v", nil,
		"Field value (required for each field)")
	itemsCreateCmd.Flags().StringArrayVar(&itemsMetaIDs, "meta-id", nil,
		"Metadata field ID (optional)")
	itemsCreateCmd.Flags().StringArrayVarP(&itemsMetaKeys, "meta-key", "K", nil,
		"Metadata field key")
	itemsCreateCmd.Flags().StringArrayVarP(&itemsMetaTypes, "meta-type", "T", nil,
		"Metadata field type")
	itemsCreateCmd.Flags().StringArrayVarP(&itemsMetaValues, "meta-value", "V", nil,
		"Metadata field value")
	_ = itemsCreateCmd.MarkFlagRequired("model")

	// Update flags
	itemsUpdateCmd.Flags().StringArrayVar(&itemsFieldIDs, "id", nil,
		"Field ID (optional, use with -k, -t, -v)")
	itemsUpdateCmd.Flags().StringArrayVarP(&itemsFieldKeys, "key", "k", nil,
		"Field key (required for each field)")
	itemsUpdateCmd.Flags().StringArrayVarP(&itemsFieldTypes, "type", "t", nil,
		"Field type (required for each field)")
	itemsUpdateCmd.Flags().StringArrayVarP(&itemsFieldValues, "value", "v", nil,
		"Field value (required for each field)")
	itemsUpdateCmd.Flags().StringArrayVar(&itemsMetaIDs, "meta-id", nil,
		"Metadata field ID (optional)")
	itemsUpdateCmd.Flags().StringArrayVarP(&itemsMetaKeys, "meta-key", "K", nil,
		"Metadata field key")
	itemsUpdateCmd.Flags().StringArrayVarP(&itemsMetaTypes, "meta-type", "T", nil,
		"Metadata field type")
	itemsUpdateCmd.Flags().StringArrayVarP(&itemsMetaValues, "meta-value", "V", nil,
		"Metadata field value")

	// Update flags - confirmation
	itemsUpdateCmd.Flags().BoolVarP(&itemsYes, "yes", "y", false,
		"Skip confirmation prompt")

	// Delete flags
	itemsDeleteCmd.Flags().BoolVarP(&itemsYes, "yes", "y", false,
		"Skip confirmation prompt")
}

func runItemsList(cmd *cobra.Command, args []string) error {
	cfg := GetConfig()
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	var items *cms.Items
	if cfg.Project != "" {
		items, err = client.GetItemsPartiallyByKey(ctx, cfg.Project,
			itemsModelID, itemsPage, itemsPerPage, itemsAsset)
	} else {
		items, err = client.GetItemsPartially(ctx, itemsModelID,
			itemsPage, itemsPerPage, itemsAsset)
	}

	if err != nil {
		return fmt.Errorf("failed to get items: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputItems(items)
}

func runItemsGet(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	item, err := client.GetItem(ctx, args[0], itemsAsset)
	if err != nil {
		return fmt.Errorf("failed to get item: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputItem(item)
}

func parseFieldsFromFlags() ([]*cms.Field, []*cms.Field, error) {
	if len(itemsFieldKeys) == 0 && len(itemsMetaKeys) == 0 {
		return nil, nil, fmt.Errorf("at least one field is required (use -k, -t, -v or -K, -T, -V)")
	}

	// Parse regular fields
	var fields []*cms.Field
	if len(itemsFieldKeys) > 0 {
		if len(itemsFieldKeys) != len(itemsFieldTypes) || len(itemsFieldKeys) != len(itemsFieldValues) {
			return nil, nil, fmt.Errorf("number of -k, -t, -v flags must match")
		}
		fields = make([]*cms.Field, len(itemsFieldKeys))
		for i := range itemsFieldKeys {
			field := &cms.Field{
				Key:   itemsFieldKeys[i],
				Type:  itemsFieldTypes[i],
				Value: itemsFieldValues[i],
			}
			if i < len(itemsFieldIDs) && itemsFieldIDs[i] != "" {
				field.ID = itemsFieldIDs[i]
			}
			fields[i] = field
		}
	}

	// Parse metadata fields
	var metaFields []*cms.Field
	if len(itemsMetaKeys) > 0 {
		if len(itemsMetaKeys) != len(itemsMetaTypes) || len(itemsMetaKeys) != len(itemsMetaValues) {
			return nil, nil, fmt.Errorf("number of -K, -T, -V flags must match")
		}
		metaFields = make([]*cms.Field, len(itemsMetaKeys))
		for i := range itemsMetaKeys {
			field := &cms.Field{
				Key:   itemsMetaKeys[i],
				Type:  itemsMetaTypes[i],
				Value: itemsMetaValues[i],
			}
			if i < len(itemsMetaIDs) && itemsMetaIDs[i] != "" {
				field.ID = itemsMetaIDs[i]
			}
			metaFields[i] = field
		}
	}

	return fields, metaFields, nil
}

func runItemsCreate(cmd *cobra.Command, args []string) error {
	cfg := GetConfig()
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	fields, metaFields, err := parseFieldsFromFlags()
	if err != nil {
		return err
	}

	ctx := context.Background()

	var item *cms.Item
	if cfg.Project != "" {
		item, err = client.CreateItemByKey(ctx, cfg.Project, itemsModelID, fields, metaFields)
	} else {
		item, err = client.CreateItem(ctx, itemsModelID, fields, metaFields)
	}
	if err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputItem(item)
}

func runItemsUpdate(cmd *cobra.Command, args []string) error {
	itemID := args[0]

	cfg := GetConfig()
	if cfg.SafeMode {
		return fmt.Errorf("update is disabled in safe mode (REEARTH_CMS_SAFE_MODE is set)")
	}

	if !itemsYes {
		if !confirmAction(fmt.Sprintf("update item %s", itemID)) {
			return nil
		}
	}

	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	fields, metaFields, err := parseFieldsFromFlags()
	if err != nil {
		return err
	}

	ctx := context.Background()
	item, err := client.UpdateItem(ctx, itemID, fields, metaFields)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputItem(item)
}

func runItemsDelete(cmd *cobra.Command, args []string) error {
	itemID := args[0]

	cfg := GetConfig()
	if cfg.SafeMode {
		return fmt.Errorf("delete is disabled in safe mode (REEARTH_CMS_SAFE_MODE is set)")
	}

	if !itemsYes {
		if !confirmAction(fmt.Sprintf("delete item %s", itemID)) {
			return nil
		}
	}

	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.DeleteItem(ctx, itemID); err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	out := NewOutputter(outputJSON)
	out.OutputMessage(fmt.Sprintf("Item %s deleted successfully", itemID))
	return nil
}

func confirmAction(action string) bool {
	fmt.Printf("Are you sure you want to %s? [y/N]: ", action)
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read response.")
		return false
	}
	response = strings.TrimSpace(strings.ToLower(response))
	if response != "y" && response != "yes" {
		fmt.Println("Aborted.")
		return false
	}
	return true
}
