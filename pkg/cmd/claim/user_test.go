package claim

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"snd-cli/pkg/cmd/claim/mocks"
	"testing"
)

func Test_addUserRun(t *testing.T) {
	tests := []struct {
		name             string
		userId           string
		claimsProvider   string
		mockResponse     string
		mockError        error
		expectedErrorMsg string
		expectedSuccess  bool
	}{
		{
			name:            "Success Case",
			userId:          "user123",
			claimsProvider:  "providerABC",
			mockResponse:    "{\"identityProvider\":\"providerABC\",\"userId\":\"user123\",\"claims\":[]}",
			mockError:       nil,
			expectedSuccess: true,
		},
		{
			name:             "Failure Case - Service Error",
			userId:           "user123",
			claimsProvider:   "providerABC",
			mockResponse:     "",
			mockError:        errors.New("service error"),
			expectedErrorMsg: "service error",
			expectedSuccess:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the Service interface
			mockService := new(mocks.Service)
			mockService.On("AddUser", tc.userId, tc.claimsProvider).Return(tc.mockResponse, tc.mockError)

			_, err := addUserRun(mockService, tc.userId, tc.claimsProvider)

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

func Test_removeUserRun(t *testing.T) {
	tests := []struct {
		name             string
		userId           string
		claimsProvider   string
		mockResponse     string
		mockError        error
		expectedErrorMsg string
		expectedSuccess  bool
	}{
		{
			name:            "Success Case",
			userId:          "user123",
			claimsProvider:  "providerABC",
			mockResponse:    "{\"identityProvider\":\"providerABC\",\"userId\":\"user123\",\"claims\":[]}",
			mockError:       nil,
			expectedSuccess: true,
		},
		{
			name:             "Failure Case - Service Error",
			userId:           "user123",
			claimsProvider:   "providerABC",
			mockResponse:     "",
			mockError:        errors.New("service error"),
			expectedErrorMsg: "service error",
			expectedSuccess:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the Service interface
			mockService := new(mocks.Service)
			mockService.On("RemoveUser", tc.userId, tc.claimsProvider).Return(tc.mockResponse, tc.mockError)

			_, err := removeUserRun(mockService, tc.userId, tc.claimsProvider)

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
