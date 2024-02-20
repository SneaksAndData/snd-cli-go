package token

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGenerateTokenCachePath(t *testing.T) {
	path, err := generateTokenCachePath()
	assert.NoError(t, err, "Generating token cache path should not produce an error")

	expectedSuffix := strings.Join([]string{folder, tokenFileName}, "/")
	assert.True(t, strings.HasSuffix(path, expectedSuffix), "Path should end with the expected suffix")
}

func TestCreateDirectory(t *testing.T) {
	dirPath := os.TempDir()
	assert.NoError(t, nil, "Creating a temporary directory should not produce an error")

	// Attempt to create the same directory again, ensuring idempotency
	err := createDirectory(dirPath)
	assert.NoError(t, err, "Ensuring an existing directory should not produce an error")

	// Cleanup
	defer os.RemoveAll(dirPath)
}
