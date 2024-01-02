package handler

import (
	"net/http"
	"strings"

	"github.com/dziablitsev/shortener/internal/response"
	"github.com/dziablitsev/shortener/internal/storage"
)

func Redirect(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimLeft(req.URL.Path, "/")
	url, found := storage.Get(id)
	if !found || req.Method != http.MethodGet {
		response.BadRequest(res)
		return
	}

	res.Header().Set("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
