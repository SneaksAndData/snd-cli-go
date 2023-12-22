package claim

import (
	"github.com/spf13/cobra"
)

func NewCmdRemoveClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a claim from an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return removeClaimRun()
		},
	}

	cmd.Flags().StringP("claims", "c", "", "Claims to add, separated by space. e.g. -c \"test1.test.sneaksanddata.com/.*:.*\" \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func removeClaimRun() error {
	// TODO: add claim remove logic, return nil if successful otherwise error
	return nil
}
