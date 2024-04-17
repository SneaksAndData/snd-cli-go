package ml

import (
	"fmt"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"snd-cli/pkg/cmdutil"
	"strings"
)

func NewCmdGetPayload(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payload",
		Short: "Get the payload for a ML Algorithm run",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("algorithm", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := getPayloadRun(http.DefaultClient, service.(*algorithmClient.Service), id, algorithm)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Specify the Crystal Job ID")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("failed to mark 'id' as a required flag: %w", err)
		return nil
	}
	return cmd
}

func getPayloadRun(client *http.Client, algorithmService Service, id, algorithm string) (string, error) {
	response, err := algorithmService.RetrievePayloadUri(id, algorithm)
	if err != nil {
		if strings.HasSuffix(err.Error(), "404") {
			return "", fmt.Errorf("failed to find run for algorithm %s with run id %s : %v", algorithm, id, "Run not found")
		}
		return "", fmt.Errorf("failed to retrieve run for algorithm %s with run id %s: %w", algorithm, id, err)
	}
	resp, err := client.Get(response.PayloadUri)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed for %s: %w", response.PayloadUri, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %w", err)
		}
		return string(respBody), nil
	case http.StatusForbidden:
		return "", fmt.Errorf("payload URI expired for algorithm %s with run id %s", algorithm, id)
	default:
		return "", fmt.Errorf("unexpected status code %d for algorithm %s with run id %s", resp.StatusCode, algorithm, id)
	}
}
