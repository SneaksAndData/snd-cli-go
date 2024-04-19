package upgrade

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
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

//go:embed scripts/install.sh
var installScript string

func installationScript() {
	updaterScriptPath, err := os.CreateTemp("", "updater-*.sh")
	if err != nil {
		fmt.Println("Error creating updater script:", err)
		os.Exit(1)
	}
	defer os.Remove(updaterScriptPath.Name())
	if _, err := updaterScriptPath.WriteString(installScript); err != nil {
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
