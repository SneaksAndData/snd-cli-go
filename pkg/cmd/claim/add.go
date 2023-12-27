package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/shared/esd-client/boxer"
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
	url := fmt.Sprintf("https://boxer-claim.%s.sneaksanddata.com", env)
	input := boxer.Input{
		TokenUrl: "",
		ClaimUrl: url,
		Auth:     boxer.ExternalToken{},
		Retries:  0,
	}
	var boxerConn boxer.Claim
	boxerConn = boxer.NewConnector(input)
	claims, err := boxerConn.AddClaim(userId, claimProvider, ca)
	if err != nil {
		return err
	}
	fmt.Println(claims)
	return nil
}
