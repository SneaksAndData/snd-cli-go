package claim

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ca []string

func NewCmdAddClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: " Add a new claim to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return addClaimRun()
		},
	}
	cmd.Flags().StringSliceVarP(&ca, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func addClaimRun() error {
	fmt.Println("Add claim")
	return nil
}
