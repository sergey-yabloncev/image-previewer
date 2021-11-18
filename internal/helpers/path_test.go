package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsExists(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "hash.go", expected: true},
		{input: "hash_exist.go", expected: false},
		{input: "", expected: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := IsExists(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
