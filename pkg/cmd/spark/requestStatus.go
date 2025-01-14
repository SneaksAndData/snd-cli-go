package spark

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

func NewCmdRequestStatus(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "request-status",
		Short:   heredoc.Doc(`Get the status of a Spark Job`),
		Example: heredoc.Doc(`snd spark request-status --id 14abbec-e517-4135-bf01-fc041a4e`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := requestStatusRun(service.(*spark.Service), id)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
	}

	return cmd

}

func requestStatusRun(sparkService Service, id string) (interface{}, error) {
	response, err := sparkService.GetLifecycleStage(id)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve lifecycle stage for run id %s: %w", id, err)
	}
	return response, nil
}
