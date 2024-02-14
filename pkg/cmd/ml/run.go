package ml

import (
	"fmt"
	"github.com/spf13/cobra"
)

var payload, tag string

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a ML Algorithm",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRun()
		},
	}

	cmd.Flags().StringVarP(&payload, "payload", "p", "", "Path to the payload JSON file")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", " Client-side submission identifier")
	return cmd
}

func runRun() error {
	url := fmt.Sprintf(crystalBaseURL, env)
	fmt.Println(url)
	panic("Not implemented")
}
