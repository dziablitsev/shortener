package handler

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRedirect(t *testing.T) {
	positiveTests := []struct {
		name     string
		url      string
		target   Target
		expected ExpectedPositive
	}{
		{
			name: "redirected successfully",
			url:  testURL,
			target: Target{
				method: http.MethodGet,
				path:   createShortURL(),
			},
			expected: ExpectedPositive{
				code: http.StatusTemporaryRedirect,
			},
		},
	}

	for _, test := range positiveTests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(test.target.method, test.target.path, nil)

			Redirect(recorder, request)
			res := recorder.Result()
			defer res.Body.Close()

			assert.Equal(t, test.expected.code, res.StatusCode)
			assert.Equal(t, test.url, res.Header.Get("Location"))
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
			target: Target{
				method: http.MethodPost,
				path:   createShortURL(),
			},
			expected: ExpectedNegative{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
				message:     "Request is not allowed!\n",
			},
		},
		{
			name: "wrong short url",
			target: Target{
				method: http.MethodGet,
				path:   createShortURL() + "wrong",
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
			request := httptest.NewRequest(test.target.method, test.target.path, nil)

			Redirect(recorder, request)
			res := recorder.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, test.expected.code, res.StatusCode)
			assert.Equal(t, test.expected.message, string(body))
		})
	}
}

func createShortURL() string {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testURL))

	Create(recorder, request)
	res := recorder.Result()
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body)
}
