package upgrade

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
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
	// Get file
	cmd = exec.Command("az", "storage", "blob", "download", "--blob-url", bundleURL, "--auth-mode", "login", "--file", fmt.Sprintf("%s/%s", basePath, arch))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error downloading binary:", err)
		os.Exit(1)
	}

	if err := os.Chmod(fmt.Sprintf("%s/%s", basePath, arch), 0755); err != nil {
		fmt.Println("Error changing file permissions:", err)
		os.Exit(1)
	}

	if _, err := os.Lstat(fmt.Sprintf("%s/.local/bin/snd", os.Getenv("HOME"))); err == nil {
		fmt.Println("Removing symlink...")
		if err := os.Remove(fmt.Sprintf("%s/.local/bin/snd", os.Getenv("HOME"))); err != nil {
			fmt.Println("Error removing symlink:", err)
			os.Exit(1)
		}
	}

	// Create a symbolic link to the application
	fmt.Println("Creating the symlink...")
	if err := os.Symlink(fmt.Sprintf("%s/%s", basePath, arch), fmt.Sprintf("%s/.local/bin/snd", os.Getenv("HOME"))); err != nil {
		fmt.Println("Error creating symlink:", err)
		os.Exit(1)
	}

	fmt.Println("Please restart your terminal for the changes to take effect.")
}
