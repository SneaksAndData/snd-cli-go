package version

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"os/exec"
	snd_cli_go "snd-cli"
	"strings"
)

const repoURL = "https://github.com/SneaksAndData/snd-cli-go"

var currentVersion = snd_cli_go.Version

func CheckIfNewVersionIsAvailable() error {
	lastTag, err := getLatestTag(repoURL)
	if err != nil {
		return fmt.Errorf("failed to get latest tag: %w", err)
	}
	result, err := compareVersions(lastTag, currentVersion)
	if err != nil {
		return fmt.Errorf("failed to compare versions: %w", err)
	}
	if result > 0 {
		fmt.Printf("New version available. Please upgrade.\nCurrent version: %s\nAvailable version: %s\nPlease run `snd upgrade` command to update the CLI to the latest version.", currentVersion, lastTag)
	} else if result < 0 {
		fmt.Printf("Your version is newer than the one present in GitHub release.\nCurrent version: %s\nLast available version in GitHub release: %s", currentVersion, lastTag)
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

func compareVersions(v1 string, v2 string) (int, error) {
	semVer1, err := semver.NewVersion(strings.TrimPrefix(v1, "v"))
	if err != nil {
		return 0, fmt.Errorf("failed to parse version %s: %w", v1, err)
	}
	semVer2, err := semver.NewVersion(strings.TrimPrefix(v2, "v"))
	if err != nil {
		return 0, fmt.Errorf("failed to parse version %s: %w", v2, err)
	}
	return semVer1.Compare(*semVer2), nil
}
