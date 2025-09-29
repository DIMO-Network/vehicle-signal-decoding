package gateways

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	sharedhttp "github.com/DIMO-Network/shared/pkg/http"

	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetDefinitionByID(t *testing.T) {
	// Define a slice to hold our test cases
	testCases := []struct {
		name             string
		mockHTTPClient   sharedhttp.ClientWrapper
		definitionID     string
		mockResponse     string
		mockResponseCode int
		expectedError    error
	}{
		{
			name: "ValidDefinitionID",
			mockHTTPClient: mockHTTPClient(func(_ *http.Request) *http.Response {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(`{ "data": { "deviceDefinition": { "model": "Some Model", "year": 2020 }}}`)),
				}
			}),
			definitionID:  "123",
			expectedError: nil,
		},
		{
			name: "InvalidDefinitionID",
			mockHTTPClient: mockHTTPClient(func(_ *http.Request) *http.Response {
				return &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(`{
  "errors": [
    {
      "message": "The is incorrect.",
      "path": [
        "deviceDefinition"
      ]
    }
  ],
  "data": null
}`)),
				}
			}),
			definitionID:  "",
			expectedError: ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := zerolog.Nop()
			settings := &config.Settings{}
			service := NewIdentityAPIService(&logger, settings, tc.mockHTTPClient)
			_, err := service.GetDefinitionByID(tc.definitionID)

			if tc.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.expectedError))
			}
		})
	}
}

type mockHTTPClient func(req *http.Request) *http.Response

// path, method string, body []byte) (*http.Response, error)
func (c mockHTTPClient) ExecuteRequest(path, method string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	// Call the mock function with the created request

	return c(req), nil
}
func (c mockHTTPClient) ExecuteRequestWithAuth(path, method string, body []byte, _ string) (*http.Response, error) {
	return c.ExecuteRequest(path, method, body)
}

func (c mockHTTPClient) GraphQLQuery(_ string, _ string, _ interface{}) error {
	return nil
}
