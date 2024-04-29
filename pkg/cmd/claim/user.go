package claim

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/spf13/cobra"
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
		Use:     "add",
		Short:   heredoc.Doc(`Add a user`),
		Example: heredoc.Doc(`snd claim user add -u user@ecco.com --claims-provider azuread`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("claim", env, url, authService)
			if err != nil {
				return err
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
		Use:     "remove",
		Short:   "Remove a user",
		Example: heredoc.Doc(`snd claim user remove -u user@ecco.com --claims-provider azuread`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("claim", env, url, authService)
			if err != nil {
				return err
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
