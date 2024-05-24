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
		payloadPath        string
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
			payloadPath:        "",
			mockFileOpErr:      nil,
			mockFileOpResponse: map[string]interface{}{"request": "request1"},
			mockResponse:       "Run successful",
			mockServiceErr:     nil,
			expectedErr:        false,
			expectedResp:       "Run successful",
		},
		{
			name:               "Failure Case - Service Error",
			algorithm:          "replenishment",
			tag:                "",
			payloadPath:        "",
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
			mockService.On("CreateRun", tc.algorithm, mock.Anything, tc.tag).Return(tc.mockResponse, tc.mockServiceErr)

			resp, err := runAlgorithm(mockService, tc.payloadPath, tc.algorithm, tc.tag)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResp, resp)
			}

			mockService.AssertExpectations(t)
		})
	}
}
