package integration_test

import (
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	MockUrl = "nginx"
	AppUrl  = "http://app:8080"
)

type AppSuite struct {
	suite.Suite
	ctx    context.Context
	client *http.Client
}

func (s *AppSuite) SetupSuite() {
	s.client = &http.Client{}
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (s *AppSuite) TestBadRequest() {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "NotImageFile",
			input: fmt.Sprintf("%s/fill/200/300/%s/test.json", AppUrl, MockUrl),
		},
		{
			name:  "NotFoundImage",
			input: fmt.Sprintf("%s/fill/200/300/%s/not-found-file.pdf", AppUrl, MockUrl),
		},
		{
			name:  "NotExistServer",
			input: fmt.Sprintf("%s/fill/200/300/%s/original_1024x504.jpg", AppUrl, "not-exist.server"),
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			req, err := http.NewRequest(http.MethodGet, tc.input, nil)
			s.Require().NoError(err)

			response, err := s.client.Do(req)
			s.Require().NoError(err)

			defer response.Body.Close()
			s.Require().NoError(err)

			s.Require().Equal(response.StatusCode, http.StatusBadRequest)
		})
	}
}

func (s *AppSuite) TestResize() {
	tests := []struct {
		with           int
		height         int
		expectedWidth  int
		expectedHeight int
	}{
		{
			with:           200,
			height:         300,
			expectedWidth:  200,
			expectedHeight: 300,
		},
		// Bigger then origin size
		{
			with:           5000,
			height:         500,
			expectedWidth:  1024,
			expectedHeight: 500,
		},
		{
			with:           1020,
			height:         5000,
			expectedWidth:  1020,
			expectedHeight: 504,
		},
	}

	for _, tc := range tests {
		s.Run(fmt.Sprintf("%vx%v", tc.with, tc.height), func() {
			req, err := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf("%s/fill/%v/%v/%s/original_1024x504.jpg", AppUrl, tc.with, tc.height, MockUrl),
				nil,
			)

			s.Require().NoError(err)

			response, err := s.client.Do(req)
			s.Require().NoError(err)

			defer response.Body.Close()
			s.Require().NoError(err)

			image, _, err := image.DecodeConfig(response.Body)

			s.Require().Equal(response.StatusCode, http.StatusOK)
			s.Require().Equal(tc.with, image.Width)
			s.Require().Equal(tc.height, image.Height)
		})
	}
}

func (s *AppSuite) TestCacheImage() {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/fill/600/400/%s/original_1024x504.jpg", AppUrl, MockUrl),
		nil,
	)

	s.Require().NoError(err)

	response, err := s.client.Do(req)
	s.Require().NoError(err)

	defer response.Body.Close()
	s.Require().NoError(err)

}
