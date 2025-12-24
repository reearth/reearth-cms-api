package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "Manage Re:Earth CMS comments",
}

var commentsItemCmd = &cobra.Command{
	Use:   "item <item-id>",
	Short: "Add a comment to an item",
	Args:  cobra.ExactArgs(1),
	RunE:  runCommentsItem,
}

var commentsAssetCmd = &cobra.Command{
	Use:   "asset <asset-id>",
	Short: "Add a comment to an asset",
	Args:  cobra.ExactArgs(1),
	RunE:  runCommentsAsset,
}

var commentsContent string

func init() {
	commentsCmd.AddCommand(commentsItemCmd)
	commentsCmd.AddCommand(commentsAssetCmd)

	commentsItemCmd.Flags().StringVarP(&commentsContent, "content", "c", "",
		"Comment content (required)")
	_ = commentsItemCmd.MarkFlagRequired("content")

	commentsAssetCmd.Flags().StringVarP(&commentsContent, "content", "c", "",
		"Comment content (required)")
	_ = commentsAssetCmd.MarkFlagRequired("content")
}

func runCommentsItem(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.CommentToItem(ctx, args[0], commentsContent); err != nil {
		return fmt.Errorf("failed to add comment to item: %w", err)
	}

	out := NewOutputter(outputJSON)
	out.OutputMessage(fmt.Sprintf("Comment added to item %s", args[0]))
	return nil
}

func runCommentsAsset(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.CommentToAsset(ctx, args[0], commentsContent); err != nil {
		return fmt.Errorf("failed to add comment to asset: %w", err)
	}

	out := NewOutputter(outputJSON)
	out.OutputMessage(fmt.Sprintf("Comment added to asset %s", args[0]))
	return nil
}
