package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_ReadJSONFile(t *testing.T) {
	tests := []struct {
		name        string
		setupFunc   func() (string, func())
		expected    map[string]interface{}
		expectedErr bool
	}{
		{
			name: "Valid JSON File",
			setupFunc: func() (string, func()) {
				content := []byte(`{"key": "value"}`)
				tmpFile, cleanup := createTempFile(t, content)
				return tmpFile, cleanup
			},
			expected:    map[string]interface{}{"key": "value"},
			expectedErr: false,
		},
		{
			name: "Invalid JSON File",
			setupFunc: func() (string, func()) {
				content := []byte(`{"key": "value"`) // Missing closing bracket
				tmpFile, cleanup := createTempFile(t, content)
				return tmpFile, cleanup
			},
			expectedErr: true,
		},
		{
			name: "File Does Not Exist",
			setupFunc: func() (string, func()) {
				return "/path/to/non/existent/file.json", func() {}
			},
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filePath, cleanup := tc.setupFunc()
			defer cleanup()

			f := File{FilePath: filePath}
			result, err := f.ReadJSONFile()

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func createTempFile(t *testing.T, content []byte) (string, func()) {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "test")
	assert.NoError(t, err)
	_, err = tmpFile.Write(content)
	assert.NoError(t, err)
	err = tmpFile.Close()
	assert.NoError(t, err)

	return tmpFile.Name(), func() { os.Remove(tmpFile.Name()) }
}

func Test_IsValidPath(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func() (string, func())
		expected  bool
	}{
		{
			name: "Existing File",
			setupFunc: func() (string, func()) {
				content := []byte(`test content`)
				tmpFile, cleanup := createTempFile(t, content)
				return tmpFile, cleanup
			},
			expected: true,
		},
		{
			name: "Non-Existing File",
			setupFunc: func() (string, func()) {
				return "/path/to/non/existent/file.json", func() {}
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filePath, cleanup := tc.setupFunc()
			defer cleanup()

			f := File{FilePath: filePath}
			result := f.IsValidPath()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_ReadAndUnmarshalWithValidFile(t *testing.T) {
	tests := []struct {
		name        string
		setupFunc   func() (string, func())
		expected    map[string]interface{}
		expectedErr bool
	}{
		{
			name: "Valid JSON File",
			setupFunc: func() (string, func()) {
				content := []byte(`{"key": "value"}`)
				tmpFile, cleanup := createTempFile(t, content)
				return tmpFile, cleanup
			},
			expected:    map[string]interface{}{"key": "value"},
			expectedErr: false,
		},
		{
			name: "Invalid JSON File",
			setupFunc: func() (string, func()) {
				content := []byte(`{"key": "value"`) // Missing closing bracket
				tmpFile, cleanup := createTempFile(t, content)
				return tmpFile, cleanup
			},
			expectedErr: true,
		},
		{
			name: "File Does Not Exist",
			setupFunc: func() (string, func()) {
				return "/path/to/non/existent/file.json", func() {}
			},
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filePath, cleanup := tc.setupFunc()
			defer cleanup()

			f := File{FilePath: filePath}
			var result map[string]interface{}
			err := f.ReadAndUnmarshal(&result)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func Test_ReadAndUnmarshalWithNonPointer(t *testing.T) {
	content := []byte(`{"key": "value"}`)
	tmpFile, cleanup := createTempFile(t, content)
	defer cleanup()

	f := File{FilePath: tmpFile}
	var result map[string]interface{}
	err := f.ReadAndUnmarshal(result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "non-nil pointer required for unmarshaling")
}

func Test_CreateDirectoryWithValidPath(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func() (string, func())
		expected  string
	}{
		{
			name: "Valid Directory Path",
			setupFunc: func() (string, func()) {
				tmpDir, err := os.MkdirTemp("", "test")
				assert.NoError(t, err)
				return tmpDir, func() { os.RemoveAll(tmpDir) }
			},
			expected: "",
		},
		{
			name: "Invalid Directory Path",
			setupFunc: func() (string, func()) {
				return "/invalid/path/to/dir", func() {}
			},
			expected: "failed to create directory",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dirPath, cleanup := tc.setupFunc()
			defer cleanup()

			f := File{FilePath: dirPath + "/testfile"}
			err := f.CreateDirectory()

			if tc.expected != "" {
				assert.Contains(t, err.Error(), tc.expected)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GenerateFilePathWithBaseHomeValidInputs(t *testing.T) {
	folderName := "testFolder"
	fileName := "testFile.txt"
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)

	expectedPath := filepath.Join(homeDir, folderName, fileName)
	result, err := GenerateFilePathWithBaseHome(folderName, fileName)
	assert.NoError(t, err)
	assert.Equal(t, expectedPath, result)
}

func Test_GenerateFilePathWithBaseHomeEmptyFolderName(t *testing.T) {
	fileName := "testFile.txt"
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)

	expectedPath := filepath.Join(homeDir, fileName)
	result, err := GenerateFilePathWithBaseHome("", fileName)
	assert.NoError(t, err)
	assert.Equal(t, expectedPath, result)
}

func Test_GenerateFilePathWithBaseHomeEmptyFileName(t *testing.T) {
	folderName := "testFolder"
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)

	expectedPath := filepath.Join(homeDir, folderName)
	result, err := GenerateFilePathWithBaseHome(folderName, "")
	assert.NoError(t, err)
	assert.Equal(t, expectedPath, result)
}

func Test_WriteToFileWithValidPath(t *testing.T) {
	content := "test content"
	tmpFile, cleanup := createTempFile(t, []byte{})
	defer cleanup()

	f := File{FilePath: tmpFile}
	err := f.WriteToFile(content)

	assert.NoError(t, err)
	writtenContent, err := os.ReadFile(tmpFile)
	assert.NoError(t, err)
	assert.Equal(t, content, string(writtenContent))
}

func Test_WriteToFileWithInvalidPath(t *testing.T) {
	content := "test content"
	invalidPath := "/invalid/path/to/file.txt"

	f := File{FilePath: invalidPath}
	err := f.WriteToFile(content)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create directory")
}

func Test_WriteToFileWithEmptyContent(t *testing.T) {
	tmpFile, cleanup := createTempFile(t, []byte{})
	defer cleanup()

	f := File{FilePath: tmpFile}
	err := f.WriteToFile("")

	assert.NoError(t, err)
	writtenContent, err := os.ReadFile(tmpFile)
	assert.NoError(t, err)
	assert.Equal(t, "", string(writtenContent))
}

func Test_WriteToFileWithNonExistentDirectory(t *testing.T) {
	content := "test content"
	tmpDir, err := os.MkdirTemp("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	nonExistentDir := filepath.Join(tmpDir, "nonexistent")
	filePath := filepath.Join(nonExistentDir, "file.txt")

	f := File{FilePath: filePath}
	err = f.WriteToFile(content)

	assert.NoError(t, err)
	writtenContent, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(writtenContent))
}
