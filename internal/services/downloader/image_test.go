package downloader_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
	"github.com/sergey-yabloncev/image-previewer/internal/services/downloader"
	"github.com/stretchr/testify/require"
)

const (
	DOMAIN = "example.com"
	URL    = "http://" + DOMAIN
)

func customResponder(statusCode int) httpmock.Responder {
	response := httpmock.NewBytesResponse(statusCode, nil)
	response.Header.Set("Content-Type", "image/jpeg")
	defer response.Body.Close()
	return httpmock.ResponderFromResponse(response)
}

func TestDownloadImage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(http.MethodGet, URL, customResponder(http.StatusOK))

	tmpDir, err := os.MkdirTemp(".", "temp")
	defer os.RemoveAll(tmpDir)

	require.NoError(t, err)

	image := tmpDir + "/test.jpg"

	err = downloader.DownloadImage(DOMAIN, image, nil)
	require.NoError(t, err)

	isExist, err := helpers.IsExists(image)
	require.NoError(t, err)
	require.True(t, isExist)

	files, _ := os.ReadDir(tmpDir)
	require.Equal(t, 1, len(files))
}

func TestNotImage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(http.MethodGet, URL, httpmock.NewStringResponder(http.StatusOK, ""))

	err := downloader.DownloadImage(DOMAIN, "test.jpg", nil)
	require.Error(t, err)
}

func TestBadStatusResponse(t *testing.T) {
	tests := []struct {
		name   string
		status int
	}{
		{name: "StatusBadRequest", status: http.StatusBadRequest},
		{name: "StatusInternalServerError", status: http.StatusInternalServerError},
		{name: "StatusNotFound", status: http.StatusNotFound},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			httpmock.Reset()
			httpmock.RegisterResponder(http.MethodGet, URL, customResponder(tc.status))

			err := downloader.DownloadImage(DOMAIN, "test.jpg", nil)
			require.Error(t, err)
		})
	}
}
