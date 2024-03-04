// Package claim /*
package claim

import (
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

const boxerClaimURL = "https://boxer-claim.%s.sneaksanddata.com"

var env, url, authProvider, userId, claimProvider string

type Service interface {
	AddClaim(user string, provider string, claims []string) (string, error)
	GetClaim(user string, provider string) (string, error)
	RemoveClaim(user string, provider string, claims []string) (string, error)
	AddUser(user string, provider string) (string, error)
	RemoveUser(user string, provider string) (string, error)
}

type ServiceFactory func(env string) (Service, error)

func NewCmdClaim(serviceFactory cmdutil.ServiceFactory, authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "claim",
		Short:   "Manage claims",
		GroupID: "claim",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&userId, "user", "u", "", "Specify the user ID")
	cmd.PersistentFlags().StringVarP(&claimProvider, "claims-provider", "", "", "Specify the claim provider")
	cmd.PersistentFlags().StringVarP(&url, "custom-service-url", "", boxerClaimURL, "Specify the service url")

	cmd.AddCommand(NewCmdUser(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRemoveClaim(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdAddClaim(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdGetClaim(authServiceFactory, serviceFactory))

	return cmd
}
