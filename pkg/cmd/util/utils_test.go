package util

import (
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
