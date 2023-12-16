package handler

import "github.com/dziablitsev/shortener/internal/config"

const testURL string = "https://practicum.yandex.ru"

type Target struct {
	method string
	path   string
}

type ExpectedPositive struct {
	code        int
	contentType string
	urlScheme   string
	urlHost     string
}

type ExpectedNegative struct {
	code        int
	contentType string
	message     string
}

func setConfig() {
	config.Server.Addr = "localhost:8080"
	config.ShortURL.Host = "http://test.ru"
	config.ShortURL.Len = 8
	config.SetConfig(config.Server, config.ShortURL)
}
