package util

import (
	"fmt"
	"os"
	"path/filepath"
)

const folder = ".snd-cli"
const file = "user-token.json"

func CacheToken(token string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to create token cache directory: %v", err)
	}
	// Construct the directory path
	dirPath := filepath.Join(homeDir, folder)

	// Create the directory if it doesn't exist
	err = os.Mkdir(dirPath, 0755)
	if err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("failed to create token cache directory: %v", err)
	}

	// Construct the file path
	filePath := filepath.Join(dirPath, file)

	// Create and open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}

	defer func() {
		errClose := file.Close()
		if err == nil && errClose != nil {
			err = errClose
		}
	}()

	// The string to be written to the file
	data := fmt.Sprint(token)

	// Write the string to the file
	_, err = file.WriteString(data)
	if err != nil {
		return "", fmt.Errorf("failed to write to file: %v", err)
	}
	return token, nil
}

func ReadToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(homeDir, folder, file)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(data), nil
}
