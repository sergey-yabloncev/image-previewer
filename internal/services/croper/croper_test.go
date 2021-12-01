package croper_test

import (
	"image"
	_ "image/jpeg"
	"os"
	"testing"

	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/croper"
	"github.com/stretchr/testify/require"
)

const (
	TestImage = "./test-data/original_1024x504.jpg"
)

func mockCropRequest(with, height int) router.CropRequest {
	return router.CropRequest{
		Type:   "fill",
		Width:  with,
		Height: height,
		URL:    "stub",
	}
}

func TestCropImage(t *testing.T) {
	tmpDir, err := os.MkdirTemp(".", "temp")
	defer os.RemoveAll(tmpDir)
	require.NoError(t, err)

	request := mockCropRequest(600, 400)

	img, err := croper.Crop(TestImage, tmpDir, "test", request)
	require.NoError(t, err)

	file, err := os.Open(img)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	require.NoError(t, err)

	image, _, err := image.DecodeConfig(file)
	require.NoError(t, err)

	require.Equal(t, request.Width, image.Width)
	require.Equal(t, request.Height, image.Height)
}
