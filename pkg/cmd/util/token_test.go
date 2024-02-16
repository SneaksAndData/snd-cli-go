package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestTokenCacheAndRead(t *testing.T) {
	token := "test-token"

	// Test CacheToken
	cachedToken, err := CacheToken(token)
	assert.NoError(t, err, "Caching token should not produce an error")
	assert.Equal(t, token, cachedToken, "Cached token should be equal to the original token")

	// Test ReadToken
	readToken, err := ReadToken()
	assert.NoError(t, err, "Reading token should not produce an error")
	assert.Equal(t, token, readToken, "Read token should be equal to the cached token")

	// Clean up: Remove test directory
	homeDir, _ := os.UserHomeDir()
	dirPath := filepath.Join(homeDir, folder)
	defer os.RemoveAll(dirPath)
}
