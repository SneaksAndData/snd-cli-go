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
