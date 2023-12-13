package handler

import (
	"github.com/dziablitsev/shortener/internal/response"
	"github.com/dziablitsev/shortener/internal/storage"
	"net/http"
	"strings"
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
