package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Original struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Target struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func Test_ConvertStruct(t *testing.T) {
	tests := []struct {
		name          string
		original      interface{}
		target        interface{}
		expectedError bool
	}{
		{
			name:          "Same structure",
			original:      Original{Name: "John", Age: 30},
			target:        &Target{},
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ConvertStruct(test.original, test.target)
			if err != nil {
				t.Errorf("ConvertStruct() error = %v, expectedError %v", err, test.expectedError)
				return
			}

			if !test.expectedError {
				original := test.original.(Original)
				target := test.target.(*Target)
				if original.Name != target.Name || original.Age != target.Age {
					t.Errorf("ConvertStruct failed, expected %v, got %v", original, target)
				}
			}
		})

	}
}

func Test_PrettifyJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Valid JSON",
			input:         `{"name":"SomeName","age":30}`,
			expected:      "{\n  \"name\": \"SomeName\",\n  \"age\": 30\n}",
			expectedError: false,
		},
		{
			name:          "Invalid JSON",
			input:         `{"name":"SomeName","age":30`,
			expected:      "",
			expectedError: true,
		},
		{
			name:          "Empty JSON",
			input:         `{}`,
			expected:      "{}",
			expectedError: false,
		},
		{
			name:          "Nested JSON",
			input:         `{"person":{"name":"SomeName","age":30}}`,
			expected:      "{\n  \"person\": {\n    \"name\": \"SomeName\",\n    \"age\": 30\n  }\n}",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PrettifyJSON(tt.input)
			if (err != nil) != tt.expectedError {
				t.Errorf("PrettifyJSON() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if result != tt.expected {
				t.Errorf("PrettifyJSON() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_ValidClaimFormat(t *testing.T) {
	tests := []struct {
		name     string
		claim    string
		expected bool
	}{
		{
			name:     "Valid claim with specific method",
			claim:    "test1.test.sneaksanddata.com/.*:GET",
			expected: true,
		},
		{
			name:     "Valid claim with wildcard method",
			claim:    "test1.test.sneaksanddata.com/.*:.*",
			expected: true,
		},
		{
			name:     "Invalid claim missing method",
			claim:    "test1.test.sneaksanddata.com/.*:",
			expected: false,
		},
		{
			name:     "Invalid claim missing path",
			claim:    ":GET",
			expected: false,
		},
		{
			name:     "Invalid claim with spaces",
			claim:    "test1.test.sneaksanddata.com/.*:GET POST",
			expected: false,
		},
		{
			name:     "Invalid claim with special characters",
			claim:    "test1.test.sneaksanddata.com/.*:G@T",
			expected: false,
		},
		{
			name:     "Invalid claim with missing path and method",
			claim:    ":",
			expected: false,
		},
		{
			name:     "Invalid claim with only path",
			claim:    "test1.test.sneaksanddata.com/.*",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateClaim(tc.claim)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_IsProdEnv(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		expected bool
	}{
		{
			name:     "awsp environment",
			env:      "awsp",
			expected: true,
		},
		{
			name:     "production environment",
			env:      "production",
			expected: true,
		},
		{
			name:     "development environment",
			env:      "development",
			expected: false,
		},
		{
			name:     "empty string environment",
			env:      "",
			expected: false,
		},
		{
			name:     "random string environment",
			env:      "random",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsProdEnv(tt.env)
			assert.Equal(t, tt.expected, result)
		})
	}
}
