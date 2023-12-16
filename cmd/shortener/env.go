package main

import (
	"os"
)

func setEnvValues() {
	if envServerAddr := os.Getenv("SERVER_ADDRESS_"); envServerAddr != "" {
		serverAddr = envServerAddr
	}
	if envBaseURL := os.Getenv("BASE_URL_"); envBaseURL != "" {
		shortURLHost = envBaseURL
	}
}
