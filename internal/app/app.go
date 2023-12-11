package app

import (
	"github.com/dziablitsev/shortener/internal/handler"
	"github.com/dziablitsev/shortener/internal/middleware"
	"net/http"
)

func Init() {
	http.Handle(`/`, middleware.URL(http.HandlerFunc(handler.Create)))

	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}
}
