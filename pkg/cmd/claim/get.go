package claim

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdGetClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves claims assigned to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getClaimRun()
		},
	}

	return cmd

}

func getClaimRun() error {
	url := fmt.Sprintf(boxerClaimBaseURL, env)
	fmt.Println(url)
	return nil
}
