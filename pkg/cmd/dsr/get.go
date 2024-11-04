package dsr

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/dsr"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

var email string

func NewCmdGetDsr(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   heredoc.Doc(`Retrieves occurrences of the specified email`),
		Example: heredoc.Doc(`snd dsr get --email user@ecco.com`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("dsr", env, url, authService)
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			resp, err := dsrRun(service.(*dsr.Service), email)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&email, "email", "", "", "Specify the  email")
	return cmd
}

func dsrRun(dsrService Service, email string) (string, error) {
	response, err := dsrService.GetDSRRequest(email)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve dsr request for email %s: %w", email, err)
	}
	return response, nil
}
