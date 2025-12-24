package main

import (
	"os"

	"github.com/reearth/reearth-cms-api/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
