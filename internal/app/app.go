package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/dziablitsev/shortener/internal/config"
	"github.com/dziablitsev/shortener/internal/handler"
	"github.com/dziablitsev/shortener/internal/logger"
	"github.com/dziablitsev/shortener/internal/middleware"
)

func Run() error {
	config.SetConfig(config.BuildConfig())

	if config.Server.Debug {
		fmt.Println("Running server on", config.Server.Addr)
		fmt.Println("Short link host is", config.ShortURL.Host)
		fmt.Println("Short link length is", config.ShortURL.Len)
	}

	if err := logger.Initialize(config.Server.LogLevel); err != nil {
		return err
	}

	err := http.ListenAndServe(config.Server.Addr, Router())
	if err != nil {
		return err
	}

	return nil
}

func Router() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.URLParam, middleware.ResponseLogger, middleware.RequestLogger)
	router.Post("/", handler.Create)
	router.Get("/{id}", handler.Redirect)
	return router
}
