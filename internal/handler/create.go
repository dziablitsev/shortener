package handler

import (
	"github.com/dziablitsev/shortener/internal/response"
	"github.com/dziablitsev/shortener/internal/storage"
	"github.com/dziablitsev/shortener/internal/url"
	"net/http"
)

func Create(res http.ResponseWriter, req *http.Request) {
	parsedURL := url.GetParsedURL(req)
	if parsedURL == "" || req.Method != http.MethodPost || req.URL.Path != "/" {
		response.BadRequest(res)
		return
	}

	key := storage.Add(parsedURL)
	shortURL := url.GetShortURL(req, key)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, err := res.Write([]byte(shortURL))
	if err != nil {
		panic(err)
	}
}
