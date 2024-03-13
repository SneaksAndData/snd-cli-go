package ml

import (
	"fmt"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
	"strings"
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

	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	return cmd
}

func cancelRun(algorithmService Service, algorithm, id string) (string, error) {
	response, err := algorithmService.CancelRun(algorithm, id)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to cancel run for algorithm %s with run id %s : %v", algorithm, id, "Run not found")
		}
		return "", fmt.Errorf("failed to cancel run for algorithm %s with run id %s: %w", algorithm, id, err)
	}
	return response, nil
}
