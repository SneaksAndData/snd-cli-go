package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

// File represents a file on the filesystem.
type File struct {
	FilePath string
}

// ReadJSONFile reads the JSON file from the file path.
func (f File) ReadJSONFile() (map[string]interface{}, error) {
	data, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read filePath '%s': %w", f.FilePath, err)
	}

	if len(data) == 0 {
		return nil, errors.New("filePath is empty")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return result, nil
}

func (f File) IsValidPath() bool {
	_, err := os.Stat(f.FilePath)
	return err == nil
}

// CreateDirectory creates a directory by using the specified path.
func (f File) CreateDirectory() error {
	err := os.MkdirAll(path.Dir(f.FilePath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

// WriteToFile writes data to the specified file path.
func (f File) WriteToFile(data string) error {
	if !f.IsValidPath() {
		err := f.CreateDirectory()
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(f.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

// GenerateFilePathWithBaseHome generates the full path for any file within a specified folder in the user's home directory.
// folderName: The name of the folder within the home directory.
// fileName: The name of the file within the specified folder.
func GenerateFilePathWithBaseHome(folderName, fileName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	dirPath := filepath.Join(homeDir, folderName) // Use the provided folder name
	return filepath.Join(dirPath, fileName), nil  // Use the provided file name
}

// ReadAndUnmarshal reads a JSON file from the provided path,
// marshals the content into a JSON string, and then unmarshal into the provided interface{}.
func (f File) ReadAndUnmarshal(v interface{}) error {
	if !f.IsValidPath() {
		return fmt.Errorf("invalid file path %s", f.FilePath)
	}

	content, err := f.ReadJSONFile()
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	c, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("error marshaling content from file: %w", err)
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("non-nil pointer required for unmarshaling")
	}

	err = json.Unmarshal(c, v)
	if err != nil {
		return fmt.Errorf("error unmarshaling content: %w", err)
	}

	return nil
}
