package storage

import (
	"math/rand"
	"time"
)

var urlMap = make(map[string]string)

const KeyLength int = 8

func Add(url string) string {
	key := generateKey(KeyLength)
	urlMap[key] = url
	return key
}

func Get(key string) (string, bool) {
	url, found := urlMap[key]
	return url, found
}

func generateKey(length int) string {
	rand.NewSource(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
