package upgrade

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func NewCmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade the CLI to the latest version",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := exec.Command("/bin/sh", "../../install.sh").Output()
			err = errors.New("this is an error")
			if err != nil {
				return err
			}
			fmt.Println(string(c))
			return nil
		},
	}
	return cmd
}
