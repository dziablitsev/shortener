package handler

import (
	"github.com/dziablitsev/shortener/internal/response"
	"github.com/dziablitsev/shortener/internal/storage"
	"net/http"
	"strings"
)

func Redirect(res http.ResponseWriter, req *http.Request) {
	key := strings.TrimLeft(req.URL.Path, "/")
	url, found := storage.Get(key)
	if !found {
		response.BadRequest(res)
		return
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
