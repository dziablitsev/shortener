package url

import (
	"github.com/dziablitsev/shortener/internal/config"
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

func GetShortURL(urlKey string) string {
	if urlKey == "" {
		return ""
	}
	return config.ShortURL.Host + "/" + urlKey
}
