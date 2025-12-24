package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Manage Re:Earth CMS models",
}

var modelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all models in a project",
	RunE:  runModelsList,
}

var modelsGetCmd = &cobra.Command{
	Use:   "get <model-id-or-key>",
	Short: "Get a specific model by ID or key",
	Args:  cobra.ExactArgs(1),
	RunE:  runModelsGet,
}

var (
	modelsProjectID string
	modelsPage      int
	modelsPerPage   int
)

func init() {
	modelsCmd.AddCommand(modelsListCmd)
	modelsCmd.AddCommand(modelsGetCmd)

	modelsListCmd.Flags().StringVarP(&modelsProjectID, "project", "p", "",
		"Project ID or alias (required)")
	modelsListCmd.Flags().IntVar(&modelsPage, "page", 1, "Page number")
	modelsListCmd.Flags().IntVar(&modelsPerPage, "per-page", 50, "Items per page")
	_ = modelsListCmd.MarkFlagRequired("project")

	modelsGetCmd.Flags().StringVarP(&modelsProjectID, "project", "p", "",
		"Project ID or alias (required when using model key)")
}

func runModelsList(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	models, err := client.GetModelsPartially(ctx, modelsProjectID, modelsPage, modelsPerPage)
	if err != nil {
		return fmt.Errorf("failed to get models: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputModels(models)
}

func runModelsGet(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	modelIDOrKey := args[0]

	// Try to get by ID first
	model, err := client.GetModel(ctx, modelIDOrKey)
	if err == nil {
		out := NewOutputter(outputJSON)
		return out.OutputModel(model)
	}

	// If failed and project is specified, try to get by key
	if modelsProjectID != "" {
		model, err = client.GetModelByKey(ctx, modelsProjectID, modelIDOrKey)
		if err == nil {
			out := NewOutputter(outputJSON)
			return out.OutputModel(model)
		}
	}

	return fmt.Errorf("failed to get model: %w", err)
}
