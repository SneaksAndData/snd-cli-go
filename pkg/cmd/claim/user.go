package claim

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdUser(factory ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage (add/remove) a user",
	}

	cmd.AddCommand(NewCmdAddUser(factory))
	cmd.AddCommand(NewCmdRemoveUser(factory))

	return cmd

}

func NewCmdAddUser(factory ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := factory(env)
			if err != nil {
				return err
			}
			resp, err := addUserRun(service, userId, claimProvider)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	return cmd

}

func addUserRun(claimService Service, userId, claimProvider string) (string, error) {
	response, err := claimService.AddUser(userId, claimProvider)
	if err != nil {
		return "", fmt.Errorf("failed to add user %s with claim provider %s: %v", userId, claimProvider, err)
	}
	return response, nil
}

func NewCmdRemoveUser(factory ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := factory(env)
			if err != nil {
				return err
			}
			resp, err := removeUserRun(service, userId, claimProvider)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	return cmd

}

func removeUserRun(claimService Service, userId, claimProvider string) (string, error) {
	response, err := claimService.RemoveUser(userId, claimProvider)
	if err != nil {
		return "", fmt.Errorf("failed to remove user %s with claim provider %s: %v", userId, claimProvider, err)
	}
	return response, nil
}
