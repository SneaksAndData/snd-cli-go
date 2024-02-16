package spark

import (
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var trimLog bool

func NewCmdLogs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Get logs from a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return logsRun(sparkService)
		},
	}

	cmd.Flags().BoolVarP(&trimLog, "trim-logs", "t", false, "Trims log to anything after STDOUT")

	return cmd
}

func logsRun(sparkService *spark.Service) error {
	response, err := sparkService.GetLogs(id)
	if err != nil {
		log.Fatalf("Failed to retrieve logs for run id %s: %v", id, err)
	}
	if trimLog {
		response = trimLogToStdout(response)
	}

	fmt.Println("Response:", response)
	return nil
}

func trimLogToStdout(logs string) string {
	logsSplit := strings.Split(logs, "\nSTDOUT:\n")
	if len(logsSplit) > 1 {
		return logsSplit[1]
	}
	return ""
}
