// Package claim /*
package claim

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

var env, authProvider, userId, claimProvider string

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
		Use:   "claim",
		Short: "Manage claims",
		Long:  "Manage user claims",
		Example: heredoc.Doc(`
			$ snd claim get --claims-provider provider --user test@ecco.com
			$ snd claim add --claims-provider provider --user test@ecco.com --claims "test1.test.sneaksanddata.com/.*:.*"
			$ snd claim remove --claims-provider provider --user test@ecco.com --claims "test1.test.sneaksanddata.com/.*:.*"

			$ snd claim user add --claims-provider provider --user test@ecco.com 
			$ snd claim user remove --claims-provider azuread --user test@ecco.com 
		`),
		GroupID: "claim",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&userId, "user", "u", "", "Specify the user ID")
	cmd.PersistentFlags().StringVarP(&claimProvider, "claims-provider", "", "", "Specify the claim provider")

	cmd.AddCommand(NewCmdUser(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdRemoveClaim(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdAddClaim(authServiceFactory, serviceFactory))
	cmd.AddCommand(NewCmdGetClaim(authServiceFactory, serviceFactory))

	return cmd
}
