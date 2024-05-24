package ml

import (
	"encoding/json"
	algorithmClient "github.com/SneaksAndData/esd-services-api-client-go/algorithm"
)

// Payload defines the structure of the request body for creating algorithm runs.
type Payload struct {
	AlgorithmParameters map[string]interface{} `validate:"required" json:"algorithm_parameters"`
	AlgorithmName       string                 `json:"algorithm_name"`
	CustomConfiguration CustomConfiguration    `json:"custom_configuration"`
	Tag                 string                 `json:"tag"`
}

func (p Payload) MarshalJSON() ([]byte, error) {
	type T struct {
		AlgorithmParameters map[string]interface{} `validate:"required" json:"AlgorithmParameters"`
		AlgorithmName       string                 `json:"AlgorithmName"`
		CustomConfiguration CustomConfiguration    `json:"CustomConfiguration"`
		Tag                 string                 `json:"Tag"`
	}

	return json.Marshal(T(p))
}

type CustomConfiguration struct {
	ImageRepository      *string              `json:"image_repository"`
	ImageTag             *string              `json:"image_tag"`
	DeadlineSeconds      *int                 `json:"deadline_seconds"`
	MaximumRetries       *int                 `json:"maximum_retries"`
	Env                  []ConfigurationEntry `json:"env"`
	Secrets              []string             `json:"secrets"`
	Args                 []ConfigurationEntry `json:"args"`
	CpuLimit             *string              `json:"cpu_limit"`
	MemoryLimit          *string              `json:"memory_limit"`
	Workgroup            *string              `json:"workgroup"`
	AdditionalWorkgroups map[string]string    `json:"additional_workgroups"`
	Version              *string              `json:"version"`
	MonitoringParameters []string             `json:"monitoring_parameters"`
	CustomResources      map[string]string    `json:"custom_resources"`
	SpeculativeAttempts  *int                 `json:"speculative_attempts"`
}

func (c CustomConfiguration) MarshalJSON() ([]byte, error) {
	type T struct {
		ImageRepository      *string              `json:"imageRepository"`
		ImageTag             *string              `json:"imageTag"`
		DeadlineSeconds      *int                 `json:"deadlineSeconds"`
		MaximumRetries       *int                 `json:"maximumRetries"`
		Env                  []ConfigurationEntry `json:"env"`
		Secrets              []string             `json:"secrets"`
		Args                 []ConfigurationEntry `json:"args"`
		CpuLimit             *string              `json:"cpuLimit"`
		MemoryLimit          *string              `json:"memoryLimit"`
		Workgroup            *string              `json:"workgroup"`
		AdditionalWorkgroups map[string]string    `json:"additionalWorkgroups"`
		Version              *string              `json:"version"`
		MonitoringParameters []string             `json:"monitoringParameters"`
		CustomResources      map[string]string    `json:"customResources"`
		SpeculativeAttempts  *int                 `json:"speculativeAttempts"`
	}

	return json.Marshal(T(c))
}

type ConfigurationEntry struct {
	Name      string                                  `json:"name"`
	Value     string                                  `json:"value"`
	ValueType *algorithmClient.ConfigurationValueType `json:"value_from"`
}

func (c ConfigurationEntry) MarshalJSON() ([]byte, error) {
	type T struct {
		Name      string                                  `json:"name"`
		Value     string                                  `json:"value"`
		ValueType *algorithmClient.ConfigurationValueType `json:"valueFrom"`
	}

	return json.Marshal(T(c))
}
