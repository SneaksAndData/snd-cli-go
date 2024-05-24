package spark

import (
	"encoding/json"
	sparkClient "github.com/SneaksAndData/esd-services-api-client-go/spark"
)

type JobParams struct {
	ClientTag           string                 `json:"client_tag"`
	ExtraArguments      map[string]interface{} `json:"extra_arguments"`
	ProjectInputs       []JobSocket            `json:"project_inputs"`
	ProjectOutputs      []JobSocket            `json:"project_outputs"`
	ExpectedParallelism *int                   `json:"expected_parallelism"`
}

func (p JobParams) MarshalJSON() ([]byte, error) {
	type T struct {
		ClientTag           string                 `json:"clientTag"`
		ExtraArguments      map[string]interface{} `json:"extraArguments"`
		ProjectInputs       []JobSocket            `json:"projectInputs"`
		ProjectOutputs      []JobSocket            `json:"projectOutputs"`
		ExpectedParallelism *int                   `json:"expectedParallelism"`
	}

	return json.Marshal(T(p))
}

type JobSocket struct {
	Alias      string `json:"alias"`
	DataPath   string `json:"data_path"`
	DataFormat string `json:"data_format"`
}

func (p JobSocket) MarshalJSON() ([]byte, error) {
	return json.Marshal(sparkClient.JobSocket(p))
}
