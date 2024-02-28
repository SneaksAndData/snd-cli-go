package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
	"snd-cli/pkg/cmdutil"
)

func NewCmdRequestStatus(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-status",
		Short: "Get the status of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				log.Fatal(err)
			}
			service, err := serviceFactory.CreateService("spark", env, authService)
			if err != nil {
				return err
			}
			resp, err := requestStatusRun(service.(*spark.Service), id)
			if err == nil {
				fmt.Println(resp)
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
