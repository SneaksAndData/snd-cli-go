package claim

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/SneaksAndData/esd-services-api-client-go/claim"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var claimsToRemove []string

func NewCmdRemoveClaim(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove",
		Short:   heredoc.Doc(`Removes a claim from an existing user`),
		Example: heredoc.Doc(`snd claim remove -c "service.test.sneaksanddata.com/.*:.*" -u user@ecco.com --claims-provider azuread`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("claim", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := removeClaimRun(service.(*claim.Service), userId, claimProvider, claimsToRemove)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringSliceVarP(&claimsToRemove, "claims", "c", []string{}, "Claims to remove. e.g. snd add -c \"test1.test.sneaksanddata.com/.*:.*\" -c \"test2.test.sneaksanddata.com/.*:.*\"")
	return cmd
}

func removeClaimRun(claimService Service, userId, claimProvider string, claimsToRemove []string) (string, error) {
	// Validate claims
	for _, c := range claimsToRemove {
		if !util.ValidateClaim(c) {
			return "", fmt.Errorf("invalid claim format: Ensure the claim string follows the pattern 'path:method'. Please review your claim string: %s", c)
		}
	}
	// Add user claims
	response, err := claimService.RemoveClaim(userId, claimProvider, claimsToRemove)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to remove claims for user %s for claim provider %s : %v", userId, claimProvider, "User not found")
		}
		return "", fmt.Errorf("failed to remove claims for user %s with claim provider %s: %w", userId, claimProvider, err)
	}

	prettifyResponse, err := util.PrettifyJSON(response)
	if err != nil {
		return "", fmt.Errorf("failed to prettify response: %w", err)
	}
	return prettifyResponse, nil
}
