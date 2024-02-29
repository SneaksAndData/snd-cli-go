package claim

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmdutil"
)

func NewCmdUser(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage (add/remove) a user",
	}

	cmd.AddCommand(NewCmdAddUser(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRemoveUser(authServiceFactory, serviceFactory))

	return cmd

}

func NewCmdAddUser(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				log.Fatal(err)
			}
			service, err := serviceFactory.CreateService("claim", env, authService)
			if err != nil {
				log.Fatal(err)
			}
			resp, err := addUserRun(service.(*claim.Service), userId, claimProvider)
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
		return "", fmt.Errorf("failed to add user %s with claim provider %s: %w", userId, claimProvider, err)
	}
	return response, nil
}

func NewCmdRemoveUser(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				log.Fatal(err)
			}
			service, err := serviceFactory.CreateService("claim", env, authService)
			if err != nil {
				log.Fatal(err)
			}
			if err != nil {
				return err
			}
			resp, err := removeUserRun(service.(*claim.Service), userId, claimProvider)
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
		return "", fmt.Errorf("failed to remove user %s with claim provider %s: %w", userId, claimProvider, err)
	}
	return response, nil
}
