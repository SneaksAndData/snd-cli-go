package ml

import (
	"fmt"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var id string

func NewCmdGet(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("algorithm", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := getRun(service.(*algorithmClient.Service), id, algorithm)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getRun(algorithmService Service, id, algorithm string) (string, error) {
	response, err := algorithmService.RetrieveRun(id, algorithm)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to find run for algorithm %s with run id %s : %v", algorithm, id, "Run not found")
		}
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %w", algorithm, id, err)
	}

	return response, nil
}
