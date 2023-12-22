package user

import "github.com/spf13/cobra"

func NewCmdAddUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return addUserRun()
		},
	}

	return cmd

}

func addUserRun() error {
	// TODO: add logic
	return nil
}
