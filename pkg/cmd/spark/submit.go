package spark

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/spf13/cobra"
	"log"
	"os"
	"snd-cli/pkg/cmd/util"
	"strings"
)

var jobName, clientTag string
var overrides string

func NewCmdSubmit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Runs the provided Beast V3 job with optional overrides",
		RunE: func(cmd *cobra.Command, args []string) error {
			var sparkService, err = InitSparkService(fmt.Sprintf(beastBaseURL, env))
			if err != nil {
				log.Fatal(err)
			}
			return submitRun(sparkService)
		},
	}

	cmd.Flags().StringVarP(&jobName, "job-name", "n", "", "Beast SparkJob or SparkJobReference resource name")
	cmd.Flags().StringVarP(&overrides, "overrides", "o", "", "Overrides for the provided job name")
	cmd.Flags().StringVarP(&clientTag, "client-tag", "t", "", "Client tag for this submission")

	return cmd
}

func submitRun(sparkService *spark.Service) error {
	params := getOverrides()
	defaultTag := generateTag()
	if clientTag == "" {
		fmt.Printf("You have not provided a client tag for this submission. Using generated tag: %s", defaultTag)
		params.ClientTag = defaultTag
	}
	params.ClientTag = clientTag
	response, err := sparkService.RunJob(params, jobName)
	if err != nil {
		log.Fatalf("Failed to submit job: %v", err)
	}

	fmt.Println("Response:", response)
	return nil
}

func getOverrides() spark.JobParams {
	if overrides == "" {
		return spark.JobParams{
			ClientTag:           "",
			ExtraArguments:      nil,
			ProjectInputs:       nil,
			ProjectOutputs:      nil,
			ExpectedParallelism: 0,
		}
	}

	if util.IsValidPath(overrides) {
		content, err := util.ReadJSONFile(overrides)
		if err != nil {
			log.Fatalf("Failed to read JSON file '%s': %v", overrides, err)
		}
		var params *spark.JobParams
		c, err := json.Marshal(content)
		if err != nil {
			log.Fatalf("Error marshaling content from file '%s': %v", overrides, err)
		}
		err = json.Unmarshal(c, &params)
		if err != nil {
			log.Fatalf("Error unmarshaling content to spark.JobParams: %v", err)
		}
		return *params
	}
	var params *spark.JobParams
	c, err := json.Marshal(overrides)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(c, &params)
	if err != nil {
		log.Fatal(err)
	}
	return *params
}

func generateTag() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to retrieve hostname: %v", err)
	}
	// Generate a random string of 12 characters (uppercase + digits)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("Error encountered while reading: %v", err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	randomString := string(b)

	defaultTag := fmt.Sprintf("cli-%s-%s", strings.ToLower(hostname), randomString)
	return defaultTag
}
