package handler

import (
	"github.com/dziablitsev/shortener/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	setConfig()

	positiveTests := []struct {
		name     string
		url      string
		target   Target
		expected ExpectedPositive
	}{
		{
			name: "short url created",
			url:  testURL,
			target: Target{
				method: http.MethodPost,
				path:   "http://test.ru/",
			},
			expected: ExpectedPositive{
				code:        http.StatusCreated,
				contentType: "text/plain",
				urlScheme:   "http",
				urlHost:     "test.ru",
			},
		},
	}

	for _, test := range positiveTests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(test.target.method, test.target.path, strings.NewReader(test.url))
			Create(recorder, request)

			res := recorder.Result()
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			parsedURL, _ := url.Parse(string(body))
			key := strings.TrimLeft(parsedURL.Path, "/")

			require.NoError(t, err)
			assert.Equal(t, test.expected.code, res.StatusCode)
			assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, test.expected.urlScheme, parsedURL.Scheme)
			assert.Equal(t, test.expected.urlHost, parsedURL.Host)
			assert.Len(t, key, config.ShortURL.Len)
		})
	}

	negativeTests := []struct {
		name     string
		url      string
		message  string
		target   Target
		expected ExpectedNegative
	}{
		{
			name: "wrong method",
			url:  testURL,
			target: Target{
				method: http.MethodGet,
				path:   "/",
			},
			expected: ExpectedNegative{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
				message:     "Request is not allowed!\n",
			},
		},
		{
			name: "wrong path",
			url:  testURL,
			target: Target{
				method: http.MethodPost,
				path:   "/test",
			},
			expected: ExpectedNegative{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
				message:     "Request is not allowed!\n",
			},
		},
	}

	for _, test := range negativeTests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(test.target.method, test.target.path, strings.NewReader(test.url))
			Create(recorder, request)

			res := recorder.Result()
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			message := string(body)

			require.NoError(t, err)
			assert.Equal(t, test.expected.code, res.StatusCode)
			assert.Equal(t, test.expected.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, test.expected.message, message)
		})
	}
}
