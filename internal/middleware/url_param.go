package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/dziablitsev/shortener/internal/config"
	"github.com/dziablitsev/shortener/internal/response"
)

func URLParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")
		if id != "" && len(id) != config.ShortURL.Len {
			response.BadRequest(res)
			return
		}
		next.ServeHTTP(res, req)
	})
}
