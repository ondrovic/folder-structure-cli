// cmd/create_test.go

package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestCreateStructure(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-folder-structure")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test case 1: Simple structure
	structure := map[string]interface{}{
		"folder1": map[string]interface{}{
			"file1.txt": nil,
		},
		"file2.txt": nil,
	}

	err = createStructure(tempDir, structure)
	if err != nil {
		t.Fatalf("createStructure failed: %v", err)
	}

	// Verify the created structure
	verifyStructure(t, tempDir, structure)
}

func TestCreateStructureWithNestedFolders(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-nested-folders")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test case 2: Nested structure
	structure := map[string]interface{}{
		"folder1": map[string]interface{}{
			"subfolder1": map[string]interface{}{
				"file1.txt": nil,
			},
			"subfolder2": map[string]interface{}{},
		},
		"file2.txt": nil,
	}

	err = createStructure(tempDir, structure)
	if err != nil {
		t.Fatalf("createStructure failed: %v", err)
	}

	// Verify the created structure
	verifyStructure(t, tempDir, structure)
}

func TestCreateStructureWithInvalidInput(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-invalid-input")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test case 3: Invalid structure (non-nil, non-map value)
	structure := map[string]interface{}{
		"folder1": "invalid",
	}

	err = createStructure(tempDir, structure)
	if err == nil {
		t.Fatalf("Expected an error for invalid structure, but got nil")
	}
}

func verifyStructure(t *testing.T, basePath string, structure map[string]interface{}) {
	for key, value := range structure {
		path := filepath.Join(basePath, key)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("Failed to stat %s: %v", path, err)
			continue
		}

		if value == nil {
			if info.IsDir() {
				t.Errorf("Expected %s to be a file, but it's a directory", path)
			}
		} else if subStructure, ok := value.(map[string]interface{}); ok {
			if !info.IsDir() {
				t.Errorf("Expected %s to be a directory, but it's not", path)
			}
			verifyStructure(t, path, subStructure)
		} else {
			t.Errorf("Unexpected value type for %s", path)
		}
	}
}

func TestRunCreate(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-run-create")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test JSON file
	jsonStructure := map[string]interface{}{
		"folder1": map[string]interface{}{
			"file1.txt": nil,
		},
		"file2.txt": nil,
	}
	jsonContent, _ := json.Marshal(jsonStructure)
	jsonFile := filepath.Join(tempDir, "test-structure.json")
	err = os.WriteFile(jsonFile, jsonContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test JSON file: %v", err)
	}

	// Create output directory
	outputDir := filepath.Join(tempDir, "output")
	err = os.Mkdir(outputDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	// Run the create command
	cmd := &cobra.Command{}
	args := []string{jsonFile, outputDir}
	runCreate(cmd, args)

	// Verify the created structure
	verifyStructure(t, outputDir, jsonStructure)
}
