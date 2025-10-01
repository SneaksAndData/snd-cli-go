package spark

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"snd-cli/pkg/cmd/spark/mocks"
	"snd-cli/pkg/cmd/util"
	"strings"
	"testing"
)

func TestGenerateTag(t *testing.T) {
	tag, _ := util.GenerateTag()
	fmt.Println(tag)
	parts := strings.Split(tag, "-")
	assert.Equal(t, "cli", parts[0], "The tag should start with 'cli'.")
}

func TestSubmitRun(t *testing.T) {
	testCases := []struct {
		name           string
		jobName        string
		clientTag      string
		overrides      string
		mockRunJobResp string
		mockRunJobErr  error
		expectedError  bool
		expectedResp   string
	}{
		{
			name:           "Successful Submission",
			jobName:        "test-job",
			clientTag:      "custom-tag",
			overrides:      "",
			mockRunJobResp: "Job submitted successfully",
			expectedError:  false,
			expectedResp:   "Job submitted successfully",
		},
		{
			name:          "Submission Failed",
			jobName:       "test-job",
			clientTag:     "custom-tag",
			overrides:     "",
			mockRunJobErr: errors.New("submission failed"),
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.Service)

			mockService.On("RunJob", mock.AnythingOfType("spark.JobParams"), tc.jobName).Return(tc.mockRunJobResp, tc.mockRunJobErr)

			resp, err := submitRun(mockService, tc.overrides, tc.jobName)

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
