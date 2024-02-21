package ml

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var id string

func NewCmdGet(factory ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := factory(env)
			if err != nil {
				return err
			}
			resp, err := getRun(service, id, algorithm)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	return cmd
}

func getRun(algorithmService Service, id, algorithm string) (string, error) {
	response, err := algorithmService.RetrieveRun(id, algorithm)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to find run for algorithm %s with run id %s : %v", algorithm, id, "Run not found")
		}
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %v", algorithm, id, err)
	}

	return response, nil
}
