package app

import (
	"github.com/dziablitsev/shortener/internal/handler"
	"github.com/dziablitsev/shortener/internal/middleware"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Run() {
	log.Fatal(http.ListenAndServe(":8080", Router()))
}

func Router() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.URLParam)
	router.Post("/", handler.Create)
	router.Get("/{id}", handler.Redirect)
	return router
}
