package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

var object string

func NewCmdRuntimeInfo(service Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runtime-info",
		Short: "Get the runtime info of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := runtimeInfoRun(service, id)
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
		return "", fmt.Errorf("failed to retrieve runtime info for run id %s: %v", id, err)
	}
	return response, nil
}
