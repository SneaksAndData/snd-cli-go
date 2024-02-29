package ml

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"snd-cli/pkg/cmd/ml/mocks"
	"testing"
)

func Test_getRun(t *testing.T) {
	tests := []struct {
		name             string
		runId            string
		algorithm        string
		mockResponse     string
		mockError        error
		expectedErrorMsg string
		expectedSuccess  bool
	}{
		{
			name:            "Success Case",
			runId:           "abc-123",
			algorithm:       "replenishment",
			mockResponse:    "{\"requestId\":\"abc-123\",\"status\":\"FAILED\",\"resultUri\":null,\"runErrorMessage\":\"CB000: Scheduling timeout.\"}\n",
			mockError:       nil,
			expectedSuccess: true,
		},
		{
			name:             "Failure Case - Service Error",
			runId:            "abc-123",
			algorithm:        "replenishment",
			mockResponse:     "",
			mockError:        errors.New("service error"),
			expectedErrorMsg: "service error",
			expectedSuccess:  false,
		},
		{
			name:             "Failure Case - Not Found Error",
			runId:            "abc-123",
			algorithm:        "replenishment",
			mockResponse:     "",
			mockError:        errors.New("not found 404"),
			expectedErrorMsg: "Run not found",
			expectedSuccess:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the Service interface
			mockService := new(mocks.Service)
			mockService.On("RetrieveRun", tc.runId, tc.algorithm).Return(tc.mockResponse, tc.mockError)

			_, err := getRun(mockService, tc.runId, tc.algorithm)

			if tc.expectedSuccess {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				if tc.expectedErrorMsg != "" {
					assert.Contains(t, err.Error(), tc.expectedErrorMsg)
				}
			}

			// Verify the mock was called as expected
			mockService.AssertExpectations(t)
		})
	}
}
