package upgrade

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func NewCmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the CLI to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			installationScript()
		},
	}
	return cmd
}

func installationScript() {
	// Determine target OS and architecture
	arch := runtime.GOARCH
	switch runtime.GOOS {
	case "darwin":
		if arch == "arm64" {
			arch = "snd-darwin-arm64"
		}
	case "linux":
		if arch == "arm64" {
			arch = "snd-linux-arm64"
		} else if arch == "amd64" || arch == "x86_64" || arch == "x64" {
			arch = "snd-linux-amd64"
		}
	default:
		fmt.Printf("Error: Unsupported OS type or architecture: %s %s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(1)
	}
	fmt.Printf("Target OS: %s\n", runtime.GOOS)
	fmt.Printf("Target ARCH: %s\n", arch)

	bundleURL := fmt.Sprintf("https://esddatalakeproduction.blob.core.windows.net/dist/snd-cli-go/%s", arch)

	// Define base path for the application
	basePath := fmt.Sprintf("%s/.local/snd-cli", os.Getenv("HOME"))
	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}

	// Check if Azure CLI is installed
	fmt.Println("Checking if Azure CLI is installed...")
	if _, err := exec.LookPath("az"); err != nil {
		fmt.Println("Azure CLI is not installed. Please install it from https://docs.microsoft.com/en-us/cli/azure/install-azure-cli")
		os.Exit(1)
	}

	// Login into Azure
	fmt.Println("Please log in to Azure...")
	cmd := exec.Command("az", "login")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error logging into Azure:", err)
		os.Exit(1)
	}

	fmt.Printf("Downloading the binary from %s\n", bundleURL)
	tempBinaryPath := fmt.Sprintf("%s/%s.tmp", basePath, arch)
	// Get file
	cmd = exec.Command("az", "storage", "blob", "download", "--blob-url", bundleURL, "--auth-mode", "login", "--file", tempBinaryPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error downloading binary:", err)
		os.Exit(1)
	}

	// Generate updater script
	updaterScript := fmt.Sprintf(`
		#!/bin/bash
		while kill -0 %d 2>/dev/null; do
			sleep 1
		done
		mv "%s" "%s/%s"
		chmod +x "%s/%s"
		# Remove the existing symlink (if it exists)
		if [ -L "%s/.local/bin/snd" ]; then
			rm "%s/.local/bin/snd"
		fi
		# Create a new symlink to the updated binary
		ln -s "%s/%s" "%s/.local/bin/snd"
		`, os.Getpid(), tempBinaryPath, basePath, arch, basePath, arch, os.Getenv("HOME"), os.Getenv("HOME"), basePath, arch, os.Getenv("HOME"))

	updaterScriptPath, err := os.CreateTemp("", "updater-*.sh")
	if err != nil {
		fmt.Println("Error creating updater script:", err)
		os.Exit(1)
	}
	defer os.Remove(updaterScriptPath.Name())
	if _, err := updaterScriptPath.WriteString(updaterScript); err != nil {
		fmt.Println("Error writing updater script:", err)
		os.Exit(1)
	}
	if err := updaterScriptPath.Chmod(0777); err != nil {
		fmt.Println("Error setting execute permission on updater script:", err)
		os.Exit(1)
	}
	if err := updaterScriptPath.Close(); err != nil {
		fmt.Println("Error closing updater script file:", err)
		os.Exit(1)
	}

	// Execute the updater script
	if err := exec.Command("nohup", updaterScriptPath.Name(), "&").Start(); err != nil {
		fmt.Println("Error executing updater script:", err)
		os.Exit(1)
	}

	fmt.Println("Update initiated. Please wait for it to complete.")
	syscall.Exit(0) // Use syscall.Exit to immediately exit
}
