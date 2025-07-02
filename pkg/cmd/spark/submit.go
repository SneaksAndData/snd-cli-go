package spark

import (
	"crypto/rand"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	sparkClient "github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"snd-cli/pkg/cmd/util"
	"snd-cli/pkg/cmd/util/file"
	"snd-cli/pkg/cmdutil"
	"strings"
)

var jobName, clientTag string
var overrides string

func NewCmdSubmit(authServiceFactory *cmdutil.AuthServiceFactory, serviceFactory cmdutil.ServiceFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "submit",
		Long: heredoc.Doc(`Runs the provided Beast V3 job with optional overrides

The overrides should be provided as a JSON file with the structure below.

If the 'clientTag' is not provided, a random tag will be generated.

If 'extraArguments', 'projectInputs', 'projectOutputs', or 'expectedParallelism' are not provided, the job will run with the default arguments.

<pre><code>
{
 "client_tag": "&lt;string&gt; - A tag for the client making the submission",
 "extra_arguments": "&lt;object&gt; - Any additional arguments for the job",
 "project_inputs": [{
	"alias": "&lt;string&gt;  - An alias for the input",
	"data_path": "&lt;string&gt; - The path to the input data",
	"data_format": "&lt;string&gt; - The format of the input data"
	}
		// More input objects can be added here
	],
 "project_outputs": [{
	"alias": "&lt;string&gt; - An alias for the output",
	"data_path": "&lt;string&gt; - The path where the output data should be stored",
	"data_format": "&lt;string&gt; - The format of the output data"
	}
		// More output objects can be added here
	],
 "expected_parallelism": "&lt;integer&gt; - The expected level of parallelism for the job"
}
</code></pre>
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			authService, err := cmdutil.InitializeAuthService(authUrl, env, authProvider, *authServiceFactory)
			if err != nil {
				return err
			}
			service, err := serviceFactory.CreateService("spark", env, url, authService)
			if err != nil {
				return err
			}
			resp, err := submitRun(service.(*sparkClient.Service), overrides, jobName)
			if err == nil {
				pterm.DefaultBasicText.Println(resp)
			}
			return err
		},
	}

	cmd.Flags().StringVarP(&jobName, "job-name", "n", "", "Beast SparkJob or SparkJobReference resource name")
	cmd.Flags().StringVarP(&overrides, "overrides", "o", "", "Overrides for the provided job name")
	cmd.Flags().StringVarP(&clientTag, "client-tag", "t", "", "Client tag for this submission")

	err := cmd.MarkFlagRequired("job-name")
	if err != nil {
		fmt.Println("failed to mark 'job-name' as a required flag: %w", err)
		return nil
	}

	return cmd
}

func submitRun(sparkService Service, overrides, jobName string) (string, error) {
	params, err := getOverrides(overrides)
	if err != nil {
		return "", err
	}
	defaultTag, _ := generateTag()
	if clientTag == "" && params.ClientTag == "" {
		pterm.DefaultBasicText.Println(pterm.Sprintf("You have not provided a client tag for this submission. Using generated tag: %s \n", defaultTag))
		params.ClientTag = defaultTag
	} else if clientTag != "" {
		params.ClientTag = clientTag
	}
	response, err := sparkService.RunJob(params, jobName)
	if err != nil {
		return "", err
	}
	return response, nil
}

func getOverrides(overrides string) (sparkClient.JobParams, error) {
	var dp = sparkClient.JobParams{
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

	var userPayload *JobParams
	err := f.ReadAndUnmarshal(&userPayload)
	if err != nil {
		return dp, err
	}

	var payload sparkClient.JobParams
	err = util.ConvertStruct(*userPayload, &payload)
	if err != nil {
		return dp, err
	}

	return payload, nil
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
