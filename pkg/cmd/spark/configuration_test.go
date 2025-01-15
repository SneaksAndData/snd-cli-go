package spark

import (
	"errors"
	"github.com/SneaksAndData/esd-services-api-client-go/spark"
	"github.com/stretchr/testify/assert"
	"snd-cli/pkg/cmd/spark/mocks"
	"testing"
)

func Test_configurationRun(t *testing.T) {
	testCases := []struct {
		name          string
		mockResponse  spark.SubmissionConfiguration
		mockError     error
		expectedError bool
		expectedResp  string
	}{
		{
			name: "ExistingConfig",
			mockResponse: spark.SubmissionConfiguration{
				RootPath: "ExistingConfig",
			},
			expectedResp:  "{\n  \"rootPath\": \"ExistingConfig\",\n  \"projectName\": \"\",\n  \"runnable\": \"\",\n  \"submissionDetails\": {\n    \"version\": \"\",\n    \"executionGroup\": \"\",\n    \"expectedParallelism\": 0,\n    \"flexibleDriver\": false,\n    \"additionalDriverNodeTolerations\": null,\n    \"maxRuntimeHours\": 0,\n    \"debugMode\": {\n      \"eventLogLocation\": \"\",\n      \"maxSizePerFile\": \"\"\n    },\n    \"submissionMode\": \"\",\n    \"extendedCodeMount\": false,\n    \"submissionJobTemplate\": \"\",\n    \"executorSpecTemplate\": \"\",\n    \"driverJobRetries\": 0,\n    \"defaultArguments\": null,\n    \"inputs\": null,\n    \"outputs\": null,\n    \"overwrite\": false\n  }\n}",
			expectedError: false,
		},
		{
			name:          "NonExistingConfig",
			mockError:     errors.New("configuration not found"),
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockService := new(mocks.Service)
			if tc.expectedError {
				mockService.On("GetConfiguration", tc.name).Return(spark.SubmissionConfiguration{}, tc.mockError)
			} else {
				mockService.On("GetConfiguration", tc.name).Return(tc.mockResponse, nil)
			}

			// Call the function under test
			resp, err := configurationRun(mockService, tc.name)

			// Assert the expectations
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResp, resp)
			}

			mockService.AssertExpectations(t)
		})
	}
}
