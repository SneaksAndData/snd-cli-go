package version

import (
	"fmt"
	"golang.org/x/mod/semver"
	"os/exec"
	snd "snd-cli/cmd"
	"strings"
)

const repoURL = "https://github.com/SneaksAndData/snd-cli-go"

var currentVersion = snd.Version

func CheckIfNewVersionIsAvailable() error {
	lastTag, err := getLatestTag(repoURL)
	if err != nil {
		return fmt.Errorf("failed to get latest tag: %w", err)
	}
	result := semver.Compare(lastTag, currentVersion)
	if result > 0 {
		fmt.Printf("New version available. Please upgrade.\nCurrent version: %s\nLast available version: %s\nPlease run `snd upgrade` command to update the CLI to the latest version.\n", currentVersion, lastTag)
	} else if result < 0 {
		fmt.Printf("Your version is newer than the one present in GitHub release.\nCurrent version: %s\nLast available version in GitHub release: %s\n", currentVersion, lastTag)
	} else {
		fmt.Printf("The snd version is up to date. %s", currentVersion)
	}

	return nil
}

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
