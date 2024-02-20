package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdRequestStatus(service Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-status",
		Short: "Get the status of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := requestStatusRun(service, id)
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
		return "", fmt.Errorf("failed to retrieve lifecycle stage for run id %s: %v", id, err)
	}
	return response, nil
}
