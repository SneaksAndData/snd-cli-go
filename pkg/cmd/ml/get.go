package ml

import (
	"fmt"
	"github.com/spf13/cobra"
)

var id string

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the result for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getRun()
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	return cmd
}

func getRun() error {
	url := fmt.Sprintf(crystalBaseURL, env)
	fmt.Println(url)
	panic("Not implemented")
}
