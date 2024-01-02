package storage

import (
	"math/rand"
	"time"

	"github.com/dziablitsev/shortener/internal/config"
)

var urlMap = make(map[string]string)

func Add(url string) string {
	id := generateID(config.ShortURL.Len)
	urlMap[id] = url
	return id
}

func Get(id string) (string, bool) {
	url, found := urlMap[id]
	return url, found
}

func generateID(length int) string {
	rand.NewSource(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
