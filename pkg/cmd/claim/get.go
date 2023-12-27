package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/shared/esd-client/boxer"
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
	url := fmt.Sprintf("https://boxer-claim.%s.sneaksanddata.com", env)
	input := boxer.Input{
		TokenUrl: "",
		ClaimUrl: url,
		Auth:     boxer.ExternalToken{},
		Retries:  0,
	}
	var boxerConn boxer.Claim
	boxerConn = boxer.NewConnector(input)
	claims, err := boxerConn.GetClaim(userId, claimProvider)
	if err != nil {
		return err
	}
	fmt.Println(claims)
	return nil
}
