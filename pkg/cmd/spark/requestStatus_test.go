package spark

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"snd-cli/pkg/cmd/spark/mocks"
	"testing"
)

func TestRequestStatusRun(t *testing.T) {
	testCases := []struct {
		name          string
		id            string // Assuming id is needed per test case
		mockResponse  interface{}
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful Status Fetch",
			id:            "valid-id",
			mockResponse:  "RUNNING",
			expectedError: false,
		},
		{
			name:          "Error Fetching Status",
			id:            "invalid-id",
			mockError:     fmt.Errorf("error fetching status"),
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.Service)

			mockService.On("GetLifecycleStage", mock.Anything).Return(tc.mockResponse, tc.mockError)

			resp, err := requestStatusRun(mockService, id)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockResponse, resp)
			}

			mockService.AssertExpectations(t)
		})
	}
}
