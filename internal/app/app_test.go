package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dziablitsev/shortener/internal/config"
)

func TestRouter(t *testing.T) {
	setConfig()

	ts := httptest.NewServer(Router())
	defer ts.Close()

	var createTestCases = []struct {
		url    string
		link   string
		method string
		status int
	}{
		{"/", "https://test.ru", http.MethodPost, http.StatusCreated},
		{"/", "http://test.ru", http.MethodPost, http.StatusCreated},
		{"/", "http://localhost", http.MethodPost, http.StatusCreated},
		{"/", "http://127.0.0.1", http.MethodPost, http.StatusCreated},
		{"/", "http://127.0.0.1:46439", http.MethodPost, http.StatusCreated},
		{"/", "http:/test.ru", http.MethodPost, http.StatusBadRequest},
		{"/", "http//test.ru", http.MethodPost, http.StatusBadRequest},
		{"/", "test.ru", http.MethodPost, http.StatusBadRequest},
		{"/", "https://test.ru", http.MethodGet, http.StatusMethodNotAllowed},
		{"/test", "", http.MethodGet, http.StatusBadRequest},
		{"/test", "", http.MethodPost, http.StatusMethodNotAllowed},
	}

	for _, test := range createTestCases {
		res := testRequest(t, ts, test.method, test.url, test.link)
		res.Body.Close()
		assert.Equal(t, test.status, res.StatusCode)
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method, path, bodyContent string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, strings.NewReader(bodyContent))
	require.NoError(t, err)

	res, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	require.NoError(t, err)

	return res
}

func setConfig() {
	config.Server.Addr = "localhost:8080"
	config.ShortURL.Host = "http://test.ru"
	config.ShortURL.Len = 8
	config.SetConfig(config.Server, config.ShortURL)
}
