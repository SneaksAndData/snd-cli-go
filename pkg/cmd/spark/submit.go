package spark

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/MakeNowJust/heredoc"
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
		Use: "submit",
		Short: heredoc.Doc(`Runs the provided Beast V3 job with optional overrides

The overrides should be provided as a JSON file with the structure below.

If the 'clientTag' is not provided, a random tag will be generated.

If 'extraArguments', 'projectInputs', 'projectOutputs', or 'expectedParallelism' are not provided, the job will run with the default arguments.

<pre><code>
{
 "clientTag": "<string> (optional) - A tag for the client making the submission",
 "extraArguments": "<object> (optional) - Any additional arguments for the job",
 "projectInputs": [{
	"alias": "<string> (optional) - An alias for the input",
	"dataPath": "<string> (required) - The path to the input data",
	"dataFormat": "<string> (required) - The format of the input data"
	}
		// More input objects can be added here
	],
 "projectOutputs": [{
	"alias": "<string> (optional) - An alias for the output",
	"dataPath": "<string> (required) - The path where the output data should be stored",
	"dataFormat": "<string> (required) - The format of the output data"
	}
		// More output objects can be added here
	],
 "expectedParallelism": "<integer> (optional) - The expected level of parallelism for the job"
}
</code></pre>
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(url, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, url, authService)
			if err != nil {
				return err
			}
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
		fmt.Printf("You have not provided a client tag for this submission. Using generated tag: %s \n", defaultTag)
		params.ClientTag = defaultTag
	} else {
		params.ClientTag = clientTag
	}
	response, err := sparkService.RunJob(params, jobName)
	if err != nil {
		return "", fmt.Errorf("failed to submit job: %w \n", err)
	}
	return response, nil
}

func getOverrides(overrides string) (spark.JobParams, error) {
	var dp = spark.JobParams{
		ClientTag:           "",
		ExtraArguments:      nil,
		ProjectInputs:       nil,
		ProjectOutputs:      nil,
		ExpectedParallelism: nil,
	}
	if overrides == "" {
		return dp, nil
	}
	f := file.File{FilePath: overrides}

	if f.IsValidPath() {
		content, err := f.ReadJSONFile()
		if err != nil {
			return dp, fmt.Errorf("failed to read JSON file '%s': %w", overrides, err)
		}
		var params *spark.JobParams
		c, err := json.Marshal(content)
		if err != nil {
			return dp, fmt.Errorf("error marshaling content from file '%s': %w", overrides, err)
		}
		err = json.Unmarshal(c, &params)
		if err != nil {
			return dp, fmt.Errorf("error unmarshaling content to spark.JobParams: %w", err)
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
		return "", fmt.Errorf("failed to retrieve hostname: %w", err)
	}
	// Generate a random string of 12 characters (uppercase + digits)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("error encountered while reading: %w", err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	randomString := string(b)

	defaultTag := fmt.Sprintf("cli-%s-%s", strings.ToLower(hostname), randomString)
	return defaultTag, nil
}
