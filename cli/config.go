package cli

import (
	"errors"
	"os"
)

type Config struct {
	BaseURL   string
	Token     string
	Workspace string
}

func LoadConfig() *Config {
	return &Config{
		BaseURL:   os.Getenv("REEARTH_CMS_BASE_URL"),
		Token:     os.Getenv("REEARTH_CMS_TOKEN"),
		Workspace: os.Getenv("REEARTH_CMS_WORKSPACE"),
	}
}

func (c *Config) ApplyFlags(baseURL, token, workspace string) {
	if baseURL != "" {
		c.BaseURL = baseURL
	}
	if token != "" {
		c.Token = token
	}
	if workspace != "" {
		c.Workspace = workspace
	}
}

func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return errors.New("base URL is required (set REEARTH_CMS_BASE_URL or use --base-url)")
	}
	if c.Token == "" {
		return errors.New("token is required (set REEARTH_CMS_TOKEN or use --token)")
	}
	return nil
}
