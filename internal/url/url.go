package url

import (
	"io"
	"net/http"
	"net/url"
)

func GetParsedURL(req *http.Request) string {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return ""
	}
	parsedURL, _ := url.Parse(string(body))
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return ""
	}
	return parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path
}

func GetShortURL(req *http.Request, urlKey string) string {
	if req.Host == "" || urlKey == "" {
		return ""
	}
	return "http://" + req.Host + "/" + urlKey
}
