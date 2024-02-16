package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestIsValidPathUsingTestify tests the IsValidPath function for both existing and non-existing paths using the testify package.
func TestIsValidPathUsingTestify(t *testing.T) {
	// Setup: Create a temporary file to test with
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	t.Run("Existing File", func(t *testing.T) {
		result := IsValidPath(tmpFile.Name())
		assert.True(t, result, "Expected path to be valid")
	})

	t.Run("Non-existing File", func(t *testing.T) {
		result := IsValidPath("/path/to/non/existing/file")
		assert.False(t, result, "Expected path to be invalid")
	})

	t.Run("JSON string", func(t *testing.T) {
		result := IsValidPath("{\"extra_arguments\":{\"client_tag\": \"\"}")
		assert.False(t, result, "Expected path to be invalid")
	})
}

func TestReadJSONFile(t *testing.T) {
	// Test for successful read and unmarshal
	t.Run("Success", func(t *testing.T) {
		jsonContent := `{"name": "John Doe", "age": 30}`
		filePath := createTempFileWithContent(t, jsonContent)
		defer os.Remove(filePath)

		result, err := ReadJSONFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"name": "John Doe", "age": float64(30)}, result)
	})

	// Test for file not found
	t.Run("FileNotFound", func(t *testing.T) {
		_, err := ReadJSONFile("/path/to/nonexistent/file.json")
		assert.Error(t, err)
	})

	// Test for empty file
	t.Run("EmptyFile", func(t *testing.T) {
		filePath := createTempFileWithContent(t, "")
		defer os.Remove(filePath)

		_, err := ReadJSONFile(filePath)
		assert.Equal(t, errors.New("file is empty"), err)
	})

	// Test for invalid JSON
	t.Run("InvalidJSON", func(t *testing.T) {
		filePath := createTempFileWithContent(t, "{invalidJson}")
		defer os.Remove(filePath)

		_, err := ReadJSONFile(filePath)
		assert.Error(t, err)
	})

	// Test for valid JSON
	t.Run("InvalidJSON", func(t *testing.T) {
		filePath := createTempFileWithContent(t, "{\"some-json\": \"hello json\"}")
		defer os.Remove(filePath)

		result, _ := ReadJSONFile(filePath)
		r := map[string]interface{}{
			"some-json": "hello json",
		}
		assert.Equal(t, r, result)
	})
}

// helper function to create a temp file with content
func createTempFileWithContent(t *testing.T, content string) string {
	t.Helper() // Marking the function as a test helper

	tmpFile, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFile.Name()
}
