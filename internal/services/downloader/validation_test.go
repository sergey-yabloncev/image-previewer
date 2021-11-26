package downloader

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const URL = "https://example.com"

func customResponder(status int, body []byte, contentType string) httpmock.Responder {
	response := httpmock.NewBytesResponse(status, body)
	response.Header.Set("Content-Type", contentType)
	defer response.Body.Close()
	return httpmock.ResponderFromResponse(response)
}

func TestIsExists(t *testing.T) {
	tests := []struct {
		contentType string
		input       string
		expected    bool
	}{
		{contentType: "application/json", input: "application/json", expected: true},
		{contentType: "application/json", input: "application/xml", expected: false},
		{contentType: "", input: "", expected: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.contentType, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder(http.MethodGet, URL, customResponder(200, nil, tc.contentType))

			request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, URL, nil)
			if err != nil {
				log.Fatal(err)
			}
			response, err := http.DefaultClient.Do(request)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			require.Equal(t, tc.expected, checkExtension(response, tc.input))
		})
	}
}
