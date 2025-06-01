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

	err = createStructure(tempDir, structure, false)
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

	err = createStructure(tempDir, structure, false)
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

	err = createStructure(tempDir, structure, false)
	if err == nil {
		t.Fatalf("Expected an error for invalid structure, but got nil")
	}
}

func TestCreateStructureWithNoOverwrite(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-no-overwrite")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create initial structure
	initialStructure := map[string]interface{}{
		"existing_folder": map[string]interface{}{
			"existing_file.txt": nil,
		},
		"existing_file.txt": nil,
	}

	err = createStructure(tempDir, initialStructure, false)
	if err != nil {
		t.Fatalf("Failed to create initial structure: %v", err)
	}

	// Write some content to the existing file to verify it's not overwritten
	existingFilePath := filepath.Join(tempDir, "existing_file.txt")
	originalContent := "original content"
	err = os.WriteFile(existingFilePath, []byte(originalContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write to existing file: %v", err)
	}

	// Also write content to the file inside the existing folder
	existingFileInFolderPath := filepath.Join(tempDir, "existing_folder", "existing_file.txt")
	existingFileContent := "existing file in folder"
	err = os.WriteFile(existingFileInFolderPath, []byte(existingFileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write to existing file in folder: %v", err)
	}

	// Try to create structure again with no-overwrite flag
	newStructure := map[string]interface{}{
		"existing_folder": map[string]interface{}{
			"existing_file.txt": nil,
			"new_file.txt":      nil,
		},
		"existing_file.txt": nil,
		"new_folder": map[string]interface{}{
			"new_file.txt": nil,
		},
	}

	err = createStructure(tempDir, newStructure, true)
	if err != nil {
		t.Fatalf("createStructure with no-overwrite failed: %v", err)
	}

	// Verify existing file content is preserved
	content, err := os.ReadFile(existingFilePath)
	if err != nil {
		t.Fatalf("Failed to read existing file: %v", err)
	}
	if string(content) != originalContent {
		t.Errorf("Expected existing file content to be preserved, got %s", string(content))
	}

	// Verify existing file in folder content is preserved
	content, err = os.ReadFile(existingFileInFolderPath)
	if err != nil {
		t.Fatalf("Failed to read existing file in folder: %v", err)
	}
	if string(content) != existingFileContent {
		t.Errorf("Expected existing file in folder content to be preserved, got %s", string(content))
	}

	// Verify new files and folders are created
	newFilePath := filepath.Join(tempDir, "existing_folder", "new_file.txt")
	if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
		t.Errorf("Expected new file to be created: %s", newFilePath)
	}

	newFolderPath := filepath.Join(tempDir, "new_folder")
	if _, err := os.Stat(newFolderPath); os.IsNotExist(err) {
		t.Errorf("Expected new folder to be created: %s", newFolderPath)
	}

	newFileInNewFolderPath := filepath.Join(tempDir, "new_folder", "new_file.txt")
	if _, err := os.Stat(newFileInNewFolderPath); os.IsNotExist(err) {
		t.Errorf("Expected new file in new folder to be created: %s", newFileInNewFolderPath)
	}
}

func TestCreateStructureOverwriteDefault(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-overwrite-default")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create initial file
	existingFilePath := filepath.Join(tempDir, "test_file.txt")
	originalContent := "original content"
	err = os.WriteFile(existingFilePath, []byte(originalContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create initial file: %v", err)
	}

	// Create structure without no-overwrite (should overwrite by default)
	structure := map[string]interface{}{
		"test_file.txt": nil,
	}

	err = createStructure(tempDir, structure, false)
	if err != nil {
		t.Fatalf("createStructure failed: %v", err)
	}

	// Verify file was overwritten (should be empty now)
	content, err := os.ReadFile(existingFilePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if len(content) != 0 {
		t.Errorf("Expected file to be overwritten (empty), but got content: %s", string(content))
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

	// Reset the noOverwrite flag to false for this test
	noOverwrite = false

	// Run the create command
	cmd := &cobra.Command{}
	args := []string{jsonFile, outputDir}
	runCreate(cmd, args)

	// Verify the created structure
	verifyStructure(t, outputDir, jsonStructure)
}
