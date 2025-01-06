// Package dsr /*
package dsr

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

const dsrURL = "https://dsr-manager.%s.sneaksanddata.com"

var env, url, authProvider, authUrl string

type Service interface {
	GetDSRRequest(email string) (string, error)
}

func NewCmdDsr(serviceFactory cmdutil.ServiceFactory, authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dsr",
		Short: "Manage DSR requests",
		Long:  "Manage DSR requests",
		Example: heredoc.Doc(`
		`),
		GroupID: "dsr",
	}
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "awsp", "Target environment")
	cmd.PersistentFlags().StringVarP(&authProvider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")
	cmd.PersistentFlags().StringVarP(&url, "custom-service-url", "", dsrURL, "Specify the service url")
	cmd.PersistentFlags().StringVarP(&authUrl, "custom-auth-url", "", "", "Specify the auth service uri")

	cmd.AddCommand(NewCmdGetDsr(authServiceFactory, serviceFactory))

	return cmd
}
