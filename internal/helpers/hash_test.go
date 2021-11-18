package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMd5String(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "nas-national-prod.s3.amazonaws.com/02mar2020_alafia_sunken_is_wads_red_banded_rosp_43_tp_106-compressed_1.jpg", expected: "8638ceb222b97c581d54d01aeb912f3b"},
		// repeat previewer
		{input: "nas-national-prod.s3.amazonaws.com/02mar2020_alafia_sunken_is_wads_red_banded_rosp_43_tp_106-compressed_1.jpg", expected: "8638ceb222b97c581d54d01aeb912f3b"},
		{input: "", expected: ""},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			require.Equal(t, tc.expected, Md5String(tc.input))
		})
	}
}
