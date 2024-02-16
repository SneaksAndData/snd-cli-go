package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
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
			var claimService, err = InitClaimService(fmt.Sprintf(boxerClaimBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return addUserRun(claimService)
		},
	}

	return cmd

}

func addUserRun(claimService *claim.Service) error {
	response, err := claimService.AddUser(userId, claimProvider)
	if err != nil {
		log.Fatalf("Failed to add user %s with claim provider %s: %v", userId, claimProvider, err)
	}

	fmt.Println("Response:", response)
	return nil
}

func NewCmdRemoveUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			var claimService, err = InitClaimService(fmt.Sprintf(boxerClaimBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return removeUserRun(claimService)
		},
	}

	return cmd

}

func removeUserRun(claimService *claim.Service) error {
	response, err := claimService.RemoveUser(userId, claimProvider)
	if err != nil {
		log.Fatalf("Failed to remove user %s with claim provider %s: %v", userId, claimProvider, err)
	}

	fmt.Println("Response:", response)
	return nil
}
