package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
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

func TestIsValidPath(t *testing.T) {
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
