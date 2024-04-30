package spark

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmdutil"
)

var object string

func NewCmdRuntimeInfo(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "runtime-info",
		Short:   heredoc.Doc(`Get the runtime info of a Spark Job`),
		Example: heredoc.Doc(`snd spark runtime-info --id 14abbec-e517-4135-bf01-fc041a4e`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, url, authService)
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			resp, err := runtimeInfoRun(service.(*spark.Service), id)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&object, "object", "o", "", "Apply a filter on the returned JSON output")

	return cmd
}

func runtimeInfoRun(sparkService Service, id string) (string, error) {
	response, err := sparkService.GetRuntimeInfo(id)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve runtime info for run id %s: %w", id, err)
	}
	return response, nil
}
