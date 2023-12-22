package claim

import (
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
	// TODO: add claim get logic, return nil if successful otherwise error
	return nil
}
