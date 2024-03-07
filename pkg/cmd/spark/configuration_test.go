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
			expectedResp:  "{\"rootPath\":\"ExistingConfig\",\"projectName\":\"\",\"runnable\":\"\",\"submissionDetails\":{\"Version\":\"\",\"ExecutionGroup\":\"\",\"ExpectedParallelism\":0,\"FlexibleDriver\":false,\"AdditionalDiverNodeTolerations\":null,\"MaxRuntimeHours\":0,\"DebugMode\":{\"EventLogLocation\":\"\",\"MaxSizePerFile\":\"\"},\"SubmissionMode\":\"\",\"ExtendedCodeMount\":false,\"SubmissionJobTemplate\":\"\",\"ExecutorSpecTemplate\":\"\",\"DriverJobRetries\":0,\"DefaultArguments\":null,\"Inputs\":null,\"Outputs\":null,\"Overwrite\":false}}",
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
