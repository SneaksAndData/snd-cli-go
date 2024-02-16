package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadJSONFile(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %v", filePath, err)
	}

	if len(data) == 0 {
		return nil, errors.New("file is empty")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return result, nil
}

func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
