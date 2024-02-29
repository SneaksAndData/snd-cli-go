package claim

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"snd-cli/pkg/cmd/claim/mocks"
	"testing"
)

func Test_getClaimRun(t *testing.T) {
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
			mockResponse:    "{\"identityProvider\":\"providerABC\",\"userId\":\"user123\",\"claims\":[{\"test.test.sneaksanddata.com/.*\":\".*\"}]}",
			mockError:       nil,
			expectedSuccess: true,
		},
		{
			name:             "Failure Case - Service Error",
			userId:           "user123",
			claimsProvider:   "providerABC",
			mockResponse:     "",
			mockError:        errors.New("service error"),
			expectedErrorMsg: "failed to retrieve claims for user user123 for claim provider providerABC : service error",
			expectedSuccess:  false,
		},
		{
			name:             "Failure Case - Not Found Error",
			userId:           "user123",
			claimsProvider:   "providerABC",
			mockResponse:     "",
			mockError:        errors.New("not found 404"),
			expectedErrorMsg: "failed to retrieve claims for user user123 for claim provider providerABC : User not found",
			expectedSuccess:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the Service interface
			mockService := new(mocks.Service)
			mockService.On("GetClaim", tc.userId, tc.claimsProvider).Return(tc.mockResponse, tc.mockError)

			_, err := getClaimRun(mockService, tc.userId, tc.claimsProvider)

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
