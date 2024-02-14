package claim

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cr []string

func NewCmdRemoveClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a claim from an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return removeClaimRun()
		},
	}

	cmd.Flags().StringSliceVarP(&cr, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func removeClaimRun() error {
	url := fmt.Sprintf(boxerClaimBaseURL, env)
	fmt.Println(url)
	panic("Not implemented")
}
