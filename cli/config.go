package cli

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const DefaultBaseURL = "https://api.cms.reearth.io"

type Config struct {
	BaseURL   string
	Token     string
	Workspace string
	Project   string
	SafeMode  bool
}

func LoadConfig() *Config {
	// Load .env file if exists (ignore error if not found)
	_ = godotenv.Load()

	baseURL := os.Getenv("REEARTH_CMS_BASE_URL")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	safeMode := strings.ToLower(os.Getenv("REEARTH_CMS_SAFE_MODE"))

	return &Config{
		BaseURL:   baseURL,
		Token:     os.Getenv("REEARTH_CMS_TOKEN"),
		Workspace: os.Getenv("REEARTH_CMS_WORKSPACE"),
		Project:   os.Getenv("REEARTH_CMS_PROJECT"),
		SafeMode:  safeMode == "true" || safeMode == "1" || safeMode == "yes",
	}
}

func (c *Config) ApplyFlags(baseURL, token, workspace, project string) {
	if baseURL != "" {
		c.BaseURL = baseURL
	}
	if token != "" {
		c.Token = token
	}
	if workspace != "" {
		c.Workspace = workspace
	}
	if project != "" {
		c.Project = project
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
