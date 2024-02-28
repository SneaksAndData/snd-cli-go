package spark

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"os"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var jobName, clientTag string
var overrides string

func NewCmdSubmit(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Runs the provided Beast V3 job with optional overrides",
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := authServiceFactory.CreateAuthService(env, authProvider)
			service, err := serviceFactory.CreateService("spark", env, authService)
			if err != nil {
				return err
			}
			resp, err := submitRun(service.(*spark.Service), overrides, jobName)
			if err == nil {
				fmt.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&jobName, "job-name", "n", "", "Beast SparkJob or SparkJobReference resource name")
	cmd.Flags().StringVarP(&overrides, "overrides", "o", "", "Overrides for the provided job name")
	cmd.Flags().StringVarP(&clientTag, "client-tag", "t", "", "Client tag for this submission")

	return cmd
}

func submitRun(sparkService Service, overrides, jobName string) (string, error) {
	params, err := getOverrides(overrides)
	if err != nil {
		return "", err
	}
	defaultTag, _ := generateTag()
	if clientTag == "" {
		fmt.Printf("You have not provided a client tag for this submission. Using generated tag: %s", defaultTag)
		params.ClientTag = defaultTag
	}
	params.ClientTag = clientTag
	response, err := sparkService.RunJob(params, jobName)
	if err != nil {
		return "", fmt.Errorf("failed to submit job: %v", err)
	}
	return response, nil
}

func getOverrides(overrides string) (spark.JobParams, error) {
	var dp = spark.JobParams{
		ClientTag:           "",
		ExtraArguments:      nil,
		ProjectInputs:       nil,
		ProjectOutputs:      nil,
		ExpectedParallelism: 0,
	}
	if overrides == "" {
		return dp, nil
	}
	f := file.File{FilePath: overrides}

	if f.IsValidPath() {
		content, err := f.ReadJSONFile()
		if err != nil {
			return dp, fmt.Errorf("failed to read JSON file '%s': %v", overrides, err)
		}
		var params *spark.JobParams
		c, err := json.Marshal(content)
		if err != nil {
			return dp, fmt.Errorf("error marshaling content from file '%s': %v", overrides, err)
		}
		err = json.Unmarshal(c, &params)
		if err != nil {
			return dp, fmt.Errorf("error unmarshaling content to spark.JobParams: %v", err)
		}
		return *params, nil
	}
	var params *spark.JobParams
	c, err := json.Marshal(overrides)
	if err != nil {
		return dp, fmt.Errorf(err.Error())
	}
	err = json.Unmarshal(c, &params)
	if err != nil {
		return dp, fmt.Errorf(err.Error())
	}
	return *params, nil
}

func generateTag() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve hostname: %v", err)
	}
	// Generate a random string of 12 characters (uppercase + digits)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("error encountered while reading: %v", err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	randomString := string(b)

	defaultTag := fmt.Sprintf("cli-%s-%s", strings.ToLower(hostname), randomString)
	return defaultTag, nil
}
