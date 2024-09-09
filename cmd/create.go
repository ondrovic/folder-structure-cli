// cmd/create.go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [json_file_path] [output_path]",
	Short: "Create folder structure from JSON",
	Long:  `Create a folder structure based on the provided JSON file.`,
	Args:  cobra.ExactArgs(2),
	Run:   runCreate,
}

func init() {
	RootCmd.AddCommand(createCmd)
}

func runCreate(cmd *cobra.Command, args []string) {
	jsonFilePath := args[0]
	outputPath := args[1]

	// Read JSON file
	jsonContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	// Parse JSON
	var structure map[string]interface{}
	err = json.Unmarshal(jsonContent, &structure)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Create folder structure
	err = createStructure(outputPath, structure)
	if err != nil {
		fmt.Printf("Error creating folder structure: %v\n", err)
		return
	}

	fmt.Println("Folder structure created successfully.")
}

func createStructure(basePath string, structure map[string]interface{}) error {
	for key, value := range structure {
		itemPath := filepath.Join(basePath, key)
		if value == nil {
			// Create file
			_, err := os.Create(itemPath)
			if err != nil {
				return fmt.Errorf("error creating file %s: %v", itemPath, err)
			}
		} else if subStructure, ok := value.(map[string]interface{}); ok {
			// Create directory
			err := os.MkdirAll(itemPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("error creating directory %s: %v", itemPath, err)
			}
			// Recursively create its structure
			err = createStructure(itemPath, subStructure)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid structure for %s", key)
		}
	}
	return nil
}
