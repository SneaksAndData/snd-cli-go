package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"snd-cli/pkg/cmd/util/token"
	"snd-cli/pkg/cmdutil"
)

var env, provider string

func NewCmdAuth(authServiceFactory *cmdutil.AuthServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Get internal authorization token",
		GroupID: "auth",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, provider)
			if err != nil {
				return fmt.Errorf("failed to initialize auth service: %w", err)
			}
			return loginRun(authService)
		},
	}

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")

	return cmd
}

func loginRun(authService token.AuthService) error {
	tokenProvider, err := token.NewProvider(authService)
	if err != nil {
		return fmt.Errorf("unable to create token provider: %w", err)
	}
	cachedToken, err := tokenProvider.GetToken() // Fetch and cache the token.
	if err != nil {
		return fmt.Errorf("unable to get the token: %w", err)
	}
	fmt.Println("Login successful.")
	fmt.Println("Token:", cachedToken)
	return nil
}
