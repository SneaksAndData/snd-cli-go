package ml

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"snd-cli/pkg/cmd/ml/mocks"
	"testing"
)

func Test_runRun(t *testing.T) {
	tests := []struct {
		name               string
		algorithm          string
		tag                string
		mockFileOpErr      error
		mockFileOpResponse map[string]interface{}
		mockResponse       string
		mockServiceErr     error
		expectedErr        bool
		expectedResp       string
	}{
		{
			name:               "Success Case",
			algorithm:          "replenishment",
			tag:                "tag1",
			mockFileOpErr:      nil,
			mockFileOpResponse: map[string]interface{}{"request": "request1"},
			mockResponse:       "{\"requestId\":\"abc-123\",\"status\":\"FAILED\",\"resultUri\":null,\"runErrorMessage\":\"CB000: Scheduling timeout.\"}\n",
			mockServiceErr:     nil,
			expectedErr:        false,
			expectedResp:       "Run successful",
		},
		{
			name:               "Failure Case - Service Error",
			algorithm:          "replenishment",
			tag:                "",
			mockFileOpResponse: map[string]interface{}{},
			mockResponse:       "",
			mockFileOpErr:      nil,
			mockServiceErr:     errors.New("service error"),
			expectedErr:        true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the Service interface
			mockService := new(mocks.Service)
			mockOperations := new(mocks.Operations)

			if tc.mockServiceErr == nil {
				mockService.On("CreateRun", tc.algorithm, mock.Anything, tc.tag).Return(tc.mockResponse, nil)
			} else {
				mockService.On("CreateRun", tc.algorithm, mock.Anything, tc.tag).Return("", tc.mockServiceErr)
			}

			// Configure mock behavior
			if tc.mockFileOpErr == nil {
				mockOperations.On("ReadJSONFile").Return(tc.mockFileOpResponse, nil)
			} else {
				mockOperations.On("ReadJSONFile").Return(map[string]interface{}{}, tc.mockFileOpErr)
			}

			mockOperations.On("IsValidPath").Return(true, nil)

			_, err := runRun(mockService, mockOperations, tc.algorithm, tc.tag)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockOperations.AssertExpectations(t)
			mockService.AssertExpectations(t)
		})
	}
}
