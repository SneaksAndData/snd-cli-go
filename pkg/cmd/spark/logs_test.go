package spark

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"snd-cli/pkg/cmd/spark/mocks"
	"testing"
)

func TestTrimLogToStdout(t *testing.T) {
	testCases := []struct {
		name         string
		logs         string
		expectedLogs string
	}{
		{
			name:         "With STDOUT",
			logs:         "INFO: starting\nSTDOUT:\nactual logs here",
			expectedLogs: "actual logs here",
		},
		{
			name:         "No STDOUT",
			logs:         "INFO: starting\nINFO: continuing",
			expectedLogs: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			trimmed := trimLogToStdout(tc.logs)
			assert.Equal(t, tc.expectedLogs, trimmed)
		})
	}
}

func TestLogsRun(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		mockLog     string
		mockError   error
		trimLog     bool
		expectedLog string
		expectedErr bool
	}{
		{
			name:        "Successful Log Fetch Without Trim",
			id:          "some-run-id",
			mockLog:     "INFO: starting\nSTDOUT:\nactual logs here",
			trimLog:     false,
			expectedLog: "INFO: starting\nSTDOUT:\nactual logs here",
			expectedErr: false,
		},
		{
			name:        "Successful Log Fetch With Trim",
			id:          "some-run-id",
			mockLog:     "INFO: starting\nSTDOUT:\nactual logs here",
			trimLog:     true,
			expectedLog: "actual logs here",
			expectedErr: false,
		},
		{
			name:        "Error Fetching Logs",
			id:          "some-run-id",
			mockError:   fmt.Errorf("network error"),
			trimLog:     false,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.Service)
			mockService.On("GetLogs", mock.Anything).Return(tc.mockLog, tc.mockError)

			resp, err := logsRun(mockService, tc.id, tc.trimLog)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedLog, resp)
			}

			mockService.AssertExpectations(t)
		})
	}
}
