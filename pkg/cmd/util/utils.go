package util

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"regexp"
	"strings"
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

// InteractiveContinue prompts the user with a confirmation message to continue running.
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

func ValidateClaim(claim string) bool {
	// Regex pattern for valid claim strings
	const claimPattern = `^([a-zA-Z0-9.-]+/[^\s:]*):([A-Za-z.*]+)$`
	re := regexp.MustCompile(claimPattern)
	return re.MatchString(claim)
}

func GenerateTag() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve hostname: %w", err)
	}
	// Generate a random string of 12 characters (uppercase + digits)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("error encountered while reading: %w", err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	randomString := string(b)

	defaultTag := fmt.Sprintf("cli-%s-%s", strings.ToLower(hostname), randomString)
	return defaultTag, nil
}
