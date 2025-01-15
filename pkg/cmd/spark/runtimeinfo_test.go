package spark

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"snd-cli/pkg/cmd/spark/mocks"
	"testing"
)

func TestRuntimeInfoRun(t *testing.T) {
	testCases := []struct {
		name          string
		id            string
		mockResponse  string
		mockError     error
		expectedError bool
		expectedResp  string
	}{
		{
			name:          "Successful Info Fetch",
			id:            "valid-id",
			mockResponse:  `{"status":"RUNNING"}`,
			expectedError: false,
			expectedResp: `{
  "status": "RUNNING"
}`,
		},
		{
			name:          "Error Fetching Info",
			id:            "invalid-id",
			mockError:     fmt.Errorf("not found"),
			expectedError: true,
		},
		// Add more test cases as needed.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock
			mockService := new(mocks.Service)
			mockService.On("GetRuntimeInfo", tc.id).Return(tc.mockResponse, tc.mockError)

			// Execute the function under test
			resp, err := runtimeInfoRun(mockService, tc.id)

			// Validate the results
			if tc.expectedError {
				assert.Error(t, err)
				if tc.mockError != nil {
					assert.Contains(t, err.Error(), tc.mockError.Error())
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResp, resp)
			}

			mockService.AssertExpectations(t)
		})
	}
}
