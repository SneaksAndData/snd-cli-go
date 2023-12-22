// Package user /*
package user

import (
	"github.com/spf13/cobra"
)

func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage (add/remove) a user",
	}

	cmd.AddCommand(NewCmdAddUser())
	cmd.AddCommand(NewCmdRemoveUser())

	return cmd

}
