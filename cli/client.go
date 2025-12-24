package cli

import (
	"fmt"

	cms "github.com/reearth/reearth-cms-api/go"
)

func NewCMSClient() (*cms.CMS, error) {
	cfg := LoadConfig()
	cfg.ApplyFlags(cfgBaseURL, cfgToken, cfgWorkspace)

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration error: %w", err)
	}

	client, err := cms.New(cfg.BaseURL, cfg.Token)
	if err != nil {
		return nil, err
	}

	if cfg.Workspace != "" {
		client = client.WithWorkspace(cfg.Workspace)
	}

	return client, nil
}
