package user

import "github.com/spf13/cobra"

func NewCmdRemoveUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return removeUserRun()
		},
	}

	return cmd

}

func removeUserRun() error {
	// TODO: add logic
	return nil
}
