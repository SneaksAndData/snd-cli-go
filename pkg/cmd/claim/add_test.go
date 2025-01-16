package claim

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"snd-cli/pkg/cmd/claim/mocks"
	"testing"
)

func Test_addClaimRun(t *testing.T) {
	tests := []struct {
		name             string
		userId           string
		claimsProvider   string
		claims           []string
		mockResponse     string
		mockError        error
		expectedErrorMsg string
		expectedSuccess  bool
		mockCall         bool
	}{
		{
			name:            "Success Case",
			userId:          "user123",
			claimsProvider:  "providerABC",
			claims:          []string{"test1.test.sneaksanddata.com/.*:.*"},
			mockResponse:    "{\"identityProvider\":\"providerABC\",\"userId\":\"user123\",\"claims\":[{\"test.test.sneaksanddata.com/.*\":\".*\"}]}",
			mockError:       nil,
			expectedSuccess: true,
			mockCall:        true,
		},
		{
			name:             "Failure Case - Service Error",
			userId:           "user123",
			claimsProvider:   "providerABC",
			claims:           []string{"test1.test.sneaksanddata.com/.*:.*"},
			mockResponse:     "",
			mockError:        errors.New("service error"),
			expectedErrorMsg: "service error",
			expectedSuccess:  false,
			mockCall:         true,
		},
		{
			name:             "Failure Case - Not Found Error",
			userId:           "user123",
			claimsProvider:   "providerABC",
			claims:           []string{"test1.test.sneaksanddata.com/.*:.*"},
			mockResponse:     "",
			mockError:        errors.New("not found 404"),
			expectedErrorMsg: "User not found",
			expectedSuccess:  false,
			mockCall:         true,
		},
		{
			name:             "Failure Case - Invalid Claim Format",
			userId:           "user123",
			claimsProvider:   "providerABC",
			claims:           []string{"test1.test.sneaksanddata.com/.*"},
			mockResponse:     "",
			mockError:        nil,
			expectedErrorMsg: "invalid claim format",
			expectedSuccess:  false,
			mockCall:         false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the Service interface
			mockService := new(mocks.Service)
			mockService.On("AddClaim", tc.userId, tc.claimsProvider, tc.claims).Return(tc.mockResponse, tc.mockError)

			_, err := addClaimRun(mockService, tc.userId, tc.claimsProvider, tc.claims)

			if tc.expectedSuccess {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				if tc.expectedErrorMsg != "" {
					assert.Contains(t, err.Error(), tc.expectedErrorMsg)
				}
			}
			if tc.mockCall {
				// Verify the mock was called as expected
				mockService.AssertExpectations(t)
			}
		})
	}
}
