package url

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/dziablitsev/shortener/internal/config"
)

func GetParsedURL(req *http.Request) (string, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	parsedURL, _ := url.Parse(string(body))
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", errors.New("parsed URL is invalid")
	}
	return parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path, nil
}

func GetShortURL(urlKey string) (string, error) {
	if urlKey == "" {
		return "", errors.New("empty key was set for short URL")
	}
	return config.ShortURL.Host + "/" + urlKey, nil
}
