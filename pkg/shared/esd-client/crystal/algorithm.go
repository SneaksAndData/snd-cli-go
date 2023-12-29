package crystal

import (
	"fmt"
	"snd-cli/pkg/shared/api"
)

type AlgorithmRequest struct {
	AlgorithmParameters interface{}
	AlgorithmName       string
	CustomConfiguration AlgorithmConfiguration
	Tag                 string
}

type AlgorithmConfiguration struct {
}

type AlgorithmRunResult struct {
	Cause   string
	Message string
	SasUri  string
}

func (c connector) CreateRun(algorithmName string, parameters interface{}, customConfig AlgorithmConfiguration, tag string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/algorithm/%s/run/%s", c.schedulerUrl, c.apiVersion, algorithmName)
	body := AlgorithmRequest{
		AlgorithmParameters: parameters,
		AlgorithmName:       algorithmName,
		CustomConfiguration: customConfig,
		Tag:                 tag,
	}
	client := api.NewClient(token)

	return client.MakeRequest("POST", targetURL, body)
}

func (c connector) RetrieveRun(runId string, algorithmName string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/algorithm/%s/results/%s/requests/%s", c.schedulerUrl, c.apiVersion, algorithmName, runId)
	client := api.NewClient(token)
	return client.MakeRequest("GET", targetURL, nil)
}

func (c connector) SubmitResult(runId string, algorithmName string, cause string, message string, sasUri string, token string) (string, error) {
	targetURL := fmt.Sprintf("%s/algorithm/%s/complete/%s/requests/%s", c.schedulerUrl, c.apiVersion, algorithmName, runId)
	fmt.Println(targetURL)
	result := AlgorithmRunResult{
		Cause:   cause,
		Message: message,
		SasUri:  sasUri,
	}
	fmt.Println(result)
	client := api.NewClient(token)

	return client.MakeRequest("POST", targetURL, result)

}
