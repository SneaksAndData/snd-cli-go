package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var ca []string

func NewCmdAddClaim(service Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new claim to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := addClaimRun(service, userId, claimProvider, ca)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringSliceVarP(&ca, "claims", "c", []string{}, "Claims to add. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func addClaimRun(claimService Service, userId, claimProvider string, ca []string) (string, error) {
	// Add user claims
	response, err := claimService.AddClaim(userId, claimProvider, ca)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to add claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to add claims for user %s with claim provider %s: %v", userId, claimProvider, err)
	}
	return response, nil
}
