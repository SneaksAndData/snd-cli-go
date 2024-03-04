package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"snd-cli/pkg/cmdutil"
	"strings"

	"github.com/spf13/cobra"
)

var trimLog bool

func NewCmdLogs(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Get logs from a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, authService)
			if err != nil {
				return err
			}
			resp, err := logsRun(service.(*spark.Service), id, trimLog)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().BoolVarP(&trimLog, "trim-logs", "t", false, "Trims log to anything after STDOUT")

	return cmd
}

func logsRun(sparkService Service, id string, trimLog bool) (string, error) {
	response, err := sparkService.GetLogs(id)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve logs for run id %s: %w", id, err)
	}
	if trimLog {
		response = trimLogToStdout(response)
	}

	return response, nil
}

func trimLogToStdout(logs string) string {
	logsSplit := strings.Split(logs, "\nSTDOUT:\n")
	if len(logsSplit) > 1 {
		return logsSplit[1]
	}
	return ""
}
