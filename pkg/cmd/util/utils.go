package util

import "encoding/json"

// ConvertStruct is a function that converts one struct to another struct.
// It takes two parameters: the original struct and the target struct.
// This function is useful when you have two structs with the same structure but different JSON keys.
func ConvertStruct(original interface{}, target interface{}) error {
	originalJSON, err := json.Marshal(original)
	if err != nil {
		return err
	}

	err = json.Unmarshal(originalJSON, target)
	if err != nil {
		return err
	}

	return nil
}
