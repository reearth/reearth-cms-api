package cli

import (
	"context"
	"encoding/json"
	"fmt"

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
	itemsModelID   string
	itemsProjectID string
	itemsPage      int
	itemsPerPage   int
	itemsAsset     bool
	itemsFields    string
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
	itemsListCmd.Flags().StringVarP(&itemsProjectID, "project", "p", "",
		"Project ID or alias (required when using model key)")
	itemsListCmd.Flags().IntVar(&itemsPage, "page", 1, "Page number")
	itemsListCmd.Flags().IntVar(&itemsPerPage, "per-page", 50, "Items per page")
	itemsListCmd.Flags().BoolVar(&itemsAsset, "asset", false, "Include asset data")
	_ = itemsListCmd.MarkFlagRequired("model")

	// Get flags
	itemsGetCmd.Flags().BoolVar(&itemsAsset, "asset", false, "Include asset data")

	// Create flags
	itemsCreateCmd.Flags().StringVarP(&itemsModelID, "model", "m", "",
		"Model ID or key (required)")
	itemsCreateCmd.Flags().StringVarP(&itemsProjectID, "project", "p", "",
		"Project ID or alias (required when using model key)")
	itemsCreateCmd.Flags().StringVarP(&itemsFields, "fields", "f", "",
		"Fields as JSON array")
	_ = itemsCreateCmd.MarkFlagRequired("model")
	_ = itemsCreateCmd.MarkFlagRequired("fields")

	// Update flags
	itemsUpdateCmd.Flags().StringVarP(&itemsFields, "fields", "f", "",
		"Fields as JSON array")
	_ = itemsUpdateCmd.MarkFlagRequired("fields")
}

func runItemsList(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	var items *cms.Items
	if itemsProjectID != "" {
		items, err = client.GetItemsPartiallyByKey(ctx, itemsProjectID,
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

func runItemsCreate(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	var fields []*cms.Field
	if err := json.Unmarshal([]byte(itemsFields), &fields); err != nil {
		return fmt.Errorf("failed to parse fields: %w", err)
	}

	ctx := context.Background()

	var item *cms.Item
	if itemsProjectID != "" {
		item, err = client.CreateItemByKey(ctx, itemsProjectID, itemsModelID, fields, nil)
	} else {
		item, err = client.CreateItem(ctx, itemsModelID, fields, nil)
	}
	if err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputItem(item)
}

func runItemsUpdate(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	var fields []*cms.Field
	if err := json.Unmarshal([]byte(itemsFields), &fields); err != nil {
		return fmt.Errorf("failed to parse fields: %w", err)
	}

	ctx := context.Background()
	item, err := client.UpdateItem(ctx, args[0], fields, nil)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputItem(item)
}

func runItemsDelete(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.DeleteItem(ctx, args[0]); err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	out := NewOutputter(outputJSON)
	out.OutputMessage(fmt.Sprintf("Item %s deleted successfully", args[0]))
	return nil
}
