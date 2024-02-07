package spark

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdRequestStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-status",
		Short: "Get the status of a Spark Job",
		RunE: func(cmd *cobra.Command, args []string) error {
			return requestStatusRun()
		},
	}

	return cmd

}

func requestStatusRun() error {
	url := fmt.Sprintf(beastBaseURL, env)
	fmt.Println(url)

	return nil
}
