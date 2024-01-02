package handler

import (
	"fmt"
	"net/http"

	"github.com/dziablitsev/shortener/internal/response"
	"github.com/dziablitsev/shortener/internal/storage"
	"github.com/dziablitsev/shortener/internal/url"
)

func Create(res http.ResponseWriter, req *http.Request) {
	parsedURL, err := url.GetParsedURL(req)
	if err != nil || req.Method != http.MethodPost || req.URL.Path != "/" {
		response.BadRequest(res)
		return
	}

	key := storage.Add(parsedURL)
	shortURL, err := url.GetShortURL(key)
	if err != nil {
		response.ShortURLError(res, fmt.Sprint(err))
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(shortURL))
	if err != nil {
		panic(err)
	}
}
