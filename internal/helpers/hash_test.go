package helpers_test

import (
	"testing"

	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestMd5String(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "compressed_1.jpg", expected: "5a670aa3b6eae0703a8ab2d1cef4b0e7d17e47cad96b95d737b9792ab43cc142"},
		// repeat previewer
		{input: "compressed_1.jpg", expected: "5a670aa3b6eae0703a8ab2d1cef4b0e7d17e47cad96b95d737b9792ab43cc142"},
		{input: "", expected: ""},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			require.Equal(t, tc.expected, helpers.Hash(tc.input))
		})
	}
}
