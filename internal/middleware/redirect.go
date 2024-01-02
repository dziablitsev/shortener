package middleware

import (
	"github.com/dziablitsev/shortener/internal/handler"
	"net/http"
)

func URL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet && req.URL.Path != "/" {
			handler.Redirect(res, req)
			return
		}
		next.ServeHTTP(res, req)
	})
}
