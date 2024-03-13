package ml

import (
	"fmt"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

func NewCmdCancel(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "Cancel a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("algorithm", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := cancelRun(service.(*algorithmClient.Service), algorithm, tag)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&tag, "tag", "t", "", "Client-side submission identifier")
	return cmd
}

func cancelRun(algorithmService Service, algorithm, tag string) (string, error) {

}
