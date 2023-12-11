package handler

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
