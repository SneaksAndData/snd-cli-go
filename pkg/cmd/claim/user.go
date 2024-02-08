package claim

import (
	"fmt"
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
	url := fmt.Sprint("https://boxer-claim.%s.sneaksanddata.com", env)
	fmt.Println(url)
	return nil
}

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
	url := fmt.Sprint(boxerClaimBaseURL, env)
	fmt.Println(url)
	return nil
}
