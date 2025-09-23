package upgrade

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func NewCmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the CLI to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			upgradeRun()
		},
	}
	return cmd
}

//go:embed scripts/upgrade.sh
var upgradeScript string

func upgradeRun() {
	updaterScriptPath, err := os.CreateTemp("", "updater-*.sh")
	if err != nil {
		fmt.Println("Error creating updater script:", err)
		os.Exit(1)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(updaterScriptPath.Name())
	if _, err := updaterScriptPath.WriteString(upgradeScript); err != nil {
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

	// Check if the OS is Linux -- if so, run the updater script in the background (does not support self updated binary)
	if strings.Contains(runtime.GOOS, "linux") {
		// Execute the updater script
		if err := exec.Command("nohup", updaterScriptPath.Name(), "&").Start(); err != nil {
			fmt.Println("Error executing updater script:", err)
			os.Exit(1)
		}

		fmt.Println("Update initiated. Please wait for it to complete.")
		syscall.Exit(0) // Use syscall.Exit to immediately exit
	}

	// Execute the updater script
	cmd := exec.Command(updaterScriptPath.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing updater script:", err)
		os.Exit(1)
	}

	fmt.Println("Update completed.")
}
