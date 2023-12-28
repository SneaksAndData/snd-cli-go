package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/boxer"
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
	url := fmt.Sprintf("https://boxer-claim.%s.sneaksanddata.com", env)
	input := boxer.Input{
		TokenUrl: "",
		ClaimUrl: url,
		Auth:     boxer.ExternalToken{},
	}
	var boxerConn boxer.User
	boxerConn = boxer.NewConnector(input)
	token, err := util.ReadToken()
	if err != nil {
		return err
	}
	user, err := boxerConn.AddUser(userId, claimProvider, token)
	if err != nil {
		return err
	}
	fmt.Println(user)
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
	url := fmt.Sprintf("https://boxer-claim.%s.sneaksanddata.com", env)
	input := boxer.Input{
		TokenUrl: "",
		ClaimUrl: url,
		Auth:     boxer.ExternalToken{},
	}
	var boxerConn boxer.User
	boxerConn = boxer.NewConnector(input)
	token, err := util.ReadToken()
	if err != nil {
		return err
	}
	_, err = boxerConn.RemoveUser(userId, claimProvider, token)
	if err != nil {
		return err
	}
	return nil
}
