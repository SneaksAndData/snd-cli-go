package claim

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func NewCmdGetClaim(service Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves claims assigned to an existing user",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := getClaimRun(service, userId, claimProvider)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	return cmd
}

func getClaimRun(claimService Service, userId, claimProvider string) (string, error) {
	// Retrieve user claims
	response, err := claimService.GetClaim(userId, claimProvider)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to retrieve claims for user %s for claim provider %s : %v", userId, claimProvider, err)
	}
	return response, nil
}
