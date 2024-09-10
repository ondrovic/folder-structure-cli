// cmd/root.go
package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "folder-structure-cli",
	Short: "Create folder structure from JSON",
	Long:  `A CLI tool to create folder structure from a JSON file.`,
}
