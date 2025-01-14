package util

import (
	"bytes"
	"encoding/json"
	"github.com/pterm/pterm"
)

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

// PrettifyJSON takes a JSON string as input and returns a pretty-printed version of the JSON string.
// It indents the JSON with two spaces for better readability.
// If there is an error during the indentation process, it returns an empty string and the error.
func PrettifyJSON(input string) (string, error) {
	var output bytes.Buffer
	err := json.Indent(&output, []byte(input), "", "  ")
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

// InteractiveContinue prompts the user with a confirmation message to continue running.i
// It restricts the user's answers to "yes" and "no".
// The function returns the user's response as a string.
func InteractiveContinue() string {
	prompt := pterm.DefaultInteractiveContinue
	prompt.DefaultText = "Are you sure you want to run in Production?"
	prompt.Options = []string{"yes", "no"}
	result, _ := prompt.Show()
	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
	return result
}

// IsProdEnv checks if the given environment string corresponds to a production environment.
// It returns true if the environment is "awsp" or "production", otherwise it returns false.
func IsProdEnv(env string) bool {
	if env == "awsp" || env == "production" {
		return true
	}
	return false
}
