package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type File struct {
	FilePath string
}

func (f File) ReadJSONFile() (map[string]interface{}, error) {
	data, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read filePath '%s': %v", f.FilePath, err)
	}

	if len(data) == 0 {
		return nil, errors.New("filePath is empty")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return result, nil
}

func (f File) IsValidPath() bool {
	_, err := os.Stat(f.FilePath)
	return err == nil
}
