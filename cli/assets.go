package cli

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	cms "github.com/reearth/reearth-cms-api/go"
	"github.com/spf13/cobra"
)

var assetsCmd = &cobra.Command{
	Use:   "asset",
	Short: "Manage Re:Earth CMS assets",
}

var assetsGetCmd = &cobra.Command{
	Use:   "get <asset-id>",
	Short: "Get asset details",
	Args:  cobra.ExactArgs(1),
	RunE:  runAssetsGet,
}

var assetsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an asset from a file (-f) or URL (-u)",
	RunE:  runAssetsCreate,
}

var assetsCatCmd = &cobra.Command{
	Use:   "cat <asset-id>",
	Short: "Output asset content to stdout",
	Args:  cobra.ExactArgs(1),
	RunE:  runAssetsCat,
}

var assetsCpCmd = &cobra.Command{
	Use:   "cp <asset-id> <destination>",
	Short: "Copy asset content to a file",
	Args:  cobra.ExactArgs(2),
	RunE:  runAssetsCp,
}

var (
	assetsFilePath string
	assetsURL      string
	assetsDirect   bool
)

func init() {
	assetsCmd.AddCommand(assetsGetCmd)
	assetsCmd.AddCommand(assetsCreateCmd)
	assetsCmd.AddCommand(assetsCatCmd)
	assetsCmd.AddCommand(assetsCpCmd)

	// Create flags
	assetsCreateCmd.Flags().StringVarP(&assetsFilePath, "file", "f", "",
		"File path to upload")
	assetsCreateCmd.Flags().StringVarP(&assetsURL, "url", "u", "",
		"URL to upload from")
	assetsCreateCmd.Flags().BoolVar(&assetsDirect, "direct", false,
		"Use direct upload instead of signed URL upload (only with -f)")
	assetsCreateCmd.MarkFlagsMutuallyExclusive("file", "url")
}

func runAssetsGet(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	asset, err := client.Asset(ctx, args[0])
	if err != nil {
		return fmt.Errorf("failed to get asset: %w", err)
	}

	out := NewOutputter(outputJSON)
	return out.OutputAsset(asset)
}

func runAssetsCreate(cmd *cobra.Command, args []string) error {
	cfg := GetConfig()
	if cfg.Project == "" {
		return fmt.Errorf("project is required (use -p or set REEARTH_CMS_PROJECT)")
	}

	if assetsFilePath == "" && assetsURL == "" {
		return fmt.Errorf("either -f (file) or -u (url) is required")
	}

	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// URL upload
	if assetsURL != "" {
		assetID, err := client.UploadAsset(ctx, cfg.Project, assetsURL)
		if err != nil {
			return fmt.Errorf("failed to upload asset from URL: %w", err)
		}

		asset, err := client.Asset(ctx, assetID)
		if err != nil {
			out := NewOutputter(outputJSON)
			out.OutputMessage(fmt.Sprintf("Asset uploaded successfully: %s", assetID))
			return nil
		}

		out := NewOutputter(outputJSON)
		return out.OutputAsset(asset)
	}

	// File upload
	file, err := os.Open(assetsFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() { _ = file.Close() }()

	var asset *cms.Asset
	if assetsDirect {
		// Direct upload
		assetID, err := client.UploadAssetDirectly(ctx, cfg.Project, assetsFilePath, file)
		if err != nil {
			return fmt.Errorf("failed to upload asset: %w", err)
		}
		asset, err = client.Asset(ctx, assetID)
		if err != nil {
			out := NewOutputter(outputJSON)
			out.OutputMessage(fmt.Sprintf("Asset uploaded successfully: %s", assetID))
			return nil
		}
	} else {
		// Signed URL upload (default)
		asset, err = uploadWithSignedURL(ctx, client, cfg.Project, assetsFilePath, file)
		if err != nil {
			return err
		}
	}

	out := NewOutputter(outputJSON)
	return out.OutputAsset(asset)
}

func uploadWithSignedURL(ctx context.Context, client *cms.CMS, projectID, filePath string, file *os.File) (*cms.Asset, error) {
	// Get file info for name
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// Step 1: Create asset upload (get signed URL)
	upload, err := client.CreateAssetUpload(ctx, projectID, info.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to create asset upload: %w", err)
	}

	// Step 2: Upload to signed URL
	if err := client.UploadToAssetUpload(ctx, upload, file); err != nil {
		return nil, fmt.Errorf("failed to upload to signed URL: %w", err)
	}

	// Step 3: Create asset by token
	asset, err := client.CreateAssetByToken(ctx, projectID, upload.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create asset: %w", err)
	}

	return asset, nil
}

func runAssetsCat(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	asset, err := client.Asset(ctx, args[0])
	if err != nil {
		return fmt.Errorf("failed to get asset: %w", err)
	}

	if asset.URL == "" {
		return fmt.Errorf("asset has no URL")
	}

	return downloadAsset(asset.URL, os.Stdout)
}

func runAssetsCp(cmd *cobra.Command, args []string) error {
	client, err := NewCMSClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	asset, err := client.Asset(ctx, args[0])
	if err != nil {
		return fmt.Errorf("failed to get asset: %w", err)
	}

	if asset.URL == "" {
		return fmt.Errorf("asset has no URL")
	}

	file, err := os.Create(args[1])
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() { _ = file.Close() }()

	if err := downloadAsset(asset.URL, file); err != nil {
		return err
	}

	out := NewOutputter(outputJSON)
	out.OutputMessage(fmt.Sprintf("Asset copied to %s", args[1]))
	return nil
}

func downloadAsset(url string, w io.Writer) error {
	resp, err := http.Get(url) //nolint:gosec,noctx
	if err != nil {
		return fmt.Errorf("failed to download asset: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download asset: status %d", resp.StatusCode)
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("failed to write asset content: %w", err)
	}

	return nil
}
