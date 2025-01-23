package example

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	cmsURL     = ""
	cmsToken   = ""
	cmsProject = ""
)

func init() {
	_ = godotenv.Load()
	cmsURL = os.Getenv("CMS_URL")
	cmsToken = os.Getenv("CMS_TOKEN")
	cmsProject = os.Getenv("CMS_PROJECT")
	if cmsURL == "" || cmsToken == "" || cmsProject == "" {
		os.Exit(0) // skip
	}
}
