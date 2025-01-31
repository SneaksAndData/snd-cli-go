package version

import (
	"fmt"
	"github.com/pterm/pterm"
	"golang.org/x/mod/semver"
	"os/exec"
	snd "snd-cli/cmd"
	"strings"
)

const repoURL = "https://github.com/SneaksAndData/snd-cli-go"

var currentVersion = snd.Version

// CheckIfNewVersionIsAvailable checks if a new version of the snd-cli is available on GitHub and prints a message to the console indicating whether an upgrade is required.
func CheckIfNewVersionIsAvailable() error {
	lastTag, err := getLatestTag(repoURL)
	if err != nil {
		return fmt.Errorf("failed to get latest tag: %w", err)
	}
	result := semver.Compare(lastTag, currentVersion)
	if result > 0 {
		pterm.DefaultBasicText.Println(
			pterm.FgLightYellow.Sprintf("New version available. Please upgrade.\n") +
				pterm.Sprintf("Current version: %s\nLast available version: %s\n", currentVersion, lastTag) +
				"Please run " +
				pterm.FgLightCyan.Sprintf("snd upgrade") +
				" command to update the CLI to the latest version.",
		)
	} else if result < 0 {
		pterm.DefaultBasicText.Println(
			"Your version is newer than the one present in GitHub release.\n" +
				pterm.Sprintf("Current version: %s\n", currentVersion) +
				pterm.Sprintf("Last available version in GitHub release: %s\n", lastTag),
		)
	} else {
		pterm.DefaultBasicText.Println(
			pterm.Sprintf("The snd version is up to date. %s\n", currentVersion),
		)
	}

	return nil
}

// getLatestTag retrieves the latest version tag for the snd-cli from GitHub.
func getLatestTag(repoURL string) (string, error) {
	// Fetch tags from the remote
	cmd := exec.Command("gh", "release", "view", "--repo", repoURL)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute gh command: %w", err)
	}
	outputLines := strings.Split(string(output), "\n")
	lastTag := ""
	if len(outputLines) > 0 {
		fields := strings.Fields(outputLines[1])
		lastTag = strings.TrimSpace(fields[1])
	}

	return lastTag, nil
}
