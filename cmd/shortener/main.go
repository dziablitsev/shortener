package main

import (
	"github.com/dziablitsev/shortener/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
