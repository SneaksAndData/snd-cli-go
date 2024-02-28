package auth

import (
	"github.com/spf13/cobra"
	"log"
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
				log.Fatalf("Failed to initialize auth service: %v", err)
			}
			return loginRun(authService)
		},
	}

	cmd.PersistentFlags().StringVarP(&env, "env", "e", "test", "Target environment")
	cmd.PersistentFlags().StringVarP(&provider, "auth-provider", "a", "azuread", "Specify the OAuth provider name")

	return cmd
}

func loginRun(authService token.AuthService) error {
	tokenProvider := token.NewProvider(authService)
	cachedToken, err := tokenProvider.GetToken() // Fetch and cache the token.
	if err != nil {
		log.Fatalf("Error logging: %v", err)
	}
	log.Println("Login successful.")
	log.Println("Token:", cachedToken)
	return nil
}
