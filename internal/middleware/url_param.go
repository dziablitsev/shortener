package middleware

import (
	"github.com/dziablitsev/shortener/internal/config"
	"github.com/dziablitsev/shortener/internal/response"
	"github.com/go-chi/chi/v5"
	"net/http"
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
