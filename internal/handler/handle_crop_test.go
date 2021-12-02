package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/sergey-yabloncev/image-previewer/internal/handler"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cache"
	"github.com/stretchr/testify/require"
)

const (
	DOMAIN = "example.com"
	URL    = "http://" + DOMAIN
)

const (
	TestImage = "./test-data/original_1024x504.jpg"
)

func customResponder() httpmock.Responder {
	response := httpmock.NewBytesResponse(http.StatusOK, httpmock.File(TestImage).Bytes())
	response.Header.Set("Content-Type", "image/jpeg")
	defer response.Body.Close()
	return httpmock.ResponderFromResponse(response)
}

func TestHandleCropImageWithCache(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(http.MethodGet, URL, customResponder())

	originImagePath, err := os.MkdirTemp(".", "originImagePath")
	defer os.RemoveAll(originImagePath)
	require.NoError(t, err)

	croppedImagePath, err := os.MkdirTemp(".", "croppedImagePath")
	defer os.RemoveAll(croppedImagePath)
	require.NoError(t, err)

	testHandler := handler.NewCropHandler(
		originImagePath,
		croppedImagePath,
		cache.NewCache(
			100,
			originImagePath,
			croppedImagePath,
			false,
		),
	)

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go func(wg *sync.WaitGroup) {
			request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/fill/200/300/"+DOMAIN, nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(testHandler.ServeHTTP)

			handler.ServeHTTP(recorder, request)

			require.Equal(t, http.StatusOK, recorder.Code)

			wg.Done()
		}(&wg)
	}
	wg.Wait()

	files, _ := os.ReadDir(originImagePath)
	require.Equal(t, 1, len(files))

	files, err = os.ReadDir(croppedImagePath)
	require.NoError(t, err)
	require.Equal(t, 1, len(files))
}
