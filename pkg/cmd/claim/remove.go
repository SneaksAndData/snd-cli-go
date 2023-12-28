package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/shared/esd-client/boxer"
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
	url := fmt.Sprintf("https://boxer-claim.%s.sneaksanddata.com", env)
	input := boxer.Input{
		TokenUrl: "",
		ClaimUrl: url,
		Auth:     boxer.ExternalToken{},
	}
	var boxerConn boxer.Claim
	boxerConn = boxer.NewConnector(input)
	token, err := util.ReadToken()
	if err != nil {
		return err
	}
	claims, err := boxerConn.RemoveClaim(userId, claimProvider, cr, token)
	if err != nil {
		return err
	}
	fmt.Println(claims)
	return nil
}
