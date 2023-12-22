package claim

import (
	"github.com/spf13/cobra"
)

func NewCmdAddClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: " Add a new claim to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return addClaimRun()
		},
	}
	cmd.Flags().StringP("claims", "c", "", "Claims to add, separated by space. e.g. -c \"test1.test.sneaksanddata.com/.*:.*\" \"test2.test.sneaksanddata.com/.*:.*\"")

	return cmd
}

func addClaimRun() error {
	// TODO: add claim add logic, return nil if successful otherwise error
	return nil
}
