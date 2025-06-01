// cmd/create.go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	noOverwrite bool
)

var createCmd = &cobra.Command{
	Use:   "create [json_file_path] [output_path]",
	Short: "Create folder structure from JSON",
	Long:  `Create a folder structure based on the provided JSON file.`,
	Args:  cobra.ExactArgs(2),
	Run:   runCreate,
}

func init() {
	createCmd.Flags().BoolVarP(&noOverwrite, "no-overwrite", "n", false, "Do not overwrite existing files and folders")
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
	err = createStructure(outputPath, structure, noOverwrite)
	if err != nil {
		fmt.Printf("Error creating folder structure: %v\n", err)
		return
	}

	fmt.Println("Folder structure created successfully.")
}

func createStructure(basePath string, structure map[string]interface{}, noOverwrite bool) error {
	for key, value := range structure {
		itemPath := filepath.Join(basePath, key)

		if value == nil {
			// Handle file creation
			if noOverwrite {
				if _, err := os.Stat(itemPath); err == nil {
					fmt.Printf("Skipping existing file: %s\n", itemPath)
					continue
				}
			}

			// Ensure the parent directory exists
			parentDir := filepath.Dir(itemPath)
			if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
				return fmt.Errorf("error creating parent directory %s: %v", parentDir, err)
			}

			_, err := os.Create(itemPath)
			if err != nil {
				return fmt.Errorf("error creating file %s: %v", itemPath, err)
			}
		} else if subStructure, ok := value.(map[string]interface{}); ok {
			// Handle directory creation
			dirExists := false
			if _, err := os.Stat(itemPath); err == nil {
				dirExists = true
				if noOverwrite {
					fmt.Printf("Directory already exists: %s\n", itemPath)
				}
			}

			// Create directory if it doesn't exist
			if !dirExists {
				err := os.MkdirAll(itemPath, os.ModePerm)
				if err != nil {
					return fmt.Errorf("error creating directory %s: %v", itemPath, err)
				}
			}

			// Always process the directory contents, regardless of whether the directory existed
			err := createStructure(itemPath, subStructure, noOverwrite)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid structure for %s", key)
		}
	}
	return nil
}
