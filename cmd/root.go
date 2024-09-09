// cmd/root.go
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version *string
)

func checkVersion() string {
	if version == nil {
		return "0.0.0"
	}

	return *version
}

var RootCmd = &cobra.Command{
	Use:     "folder-structure-cli",
	Short:   "Create folder structure from JSON",
	Long:    `A CLI tool to create folder structure from a JSON file.`,
	Version: checkVersion(),
}
