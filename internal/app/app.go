package app

import (
	"fmt"
	"github.com/dziablitsev/shortener/internal/config"
	"github.com/dziablitsev/shortener/internal/handler"
	"github.com/dziablitsev/shortener/internal/middleware"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Run() error {
	if config.Server.Debug {
		fmt.Println("Running server on", config.Server.Addr)
		fmt.Println("Short link host is", config.ShortURL.Host)
		fmt.Println("Short link length is", config.ShortURL.Len)
	}

	err := http.ListenAndServe(config.Server.Addr, Router())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func Router() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.URLParam)
	router.Post("/", handler.Create)
	router.Get("/{id}", handler.Redirect)
	return router
}
