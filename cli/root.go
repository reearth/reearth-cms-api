package cli

import (
	"github.com/spf13/cobra"
)

var (
	cfgBaseURL   string
	cfgToken     string
	cfgWorkspace string
	outputJSON   string
)

var rootCmd = &cobra.Command{
	Use:   "cms",
	Short: "Re:Earth CMS command line interface",
	Long:  "A CLI tool for interacting with Re:Earth CMS API.",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgBaseURL, "base-url", "",
		"Re:Earth CMS API base URL (default: https://api.cms.reearth.io, or set REEARTH_CMS_BASE_URL env)")
	rootCmd.PersistentFlags().StringVar(&cfgToken, "token", "",
		"API token (or set REEARTH_CMS_TOKEN env)")
	rootCmd.PersistentFlags().StringVarP(&cfgWorkspace, "workspace", "w", "",
		"Workspace ID (or set REEARTH_CMS_WORKSPACE env)")
	rootCmd.PersistentFlags().StringVar(&outputJSON, "json", "",
		"Output as JSON. Optionally specify fields: --json id,name")

	rootCmd.AddCommand(modelsCmd)
	rootCmd.AddCommand(itemsCmd)
	rootCmd.AddCommand(assetsCmd)
	rootCmd.AddCommand(commentsCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
