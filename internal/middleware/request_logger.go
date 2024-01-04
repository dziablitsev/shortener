package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/dziablitsev/shortener/internal/logger"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()

		next.ServeHTTP(res, req)

		logger.Log.Info("HTTP request",
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.Duration("duration", time.Since(start)),
		)
	})
}
