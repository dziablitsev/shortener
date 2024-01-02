package storage

import (
	"github.com/dziablitsev/shortener/internal/config"
	"math/rand"
	"time"
)

var urlMap = make(map[string]string)

func Add(url string) string {
	id := generateID(config.ShortLinkLen)
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
