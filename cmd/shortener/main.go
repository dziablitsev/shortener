package main

import (
	"github.com/dziablitsev/shortener/internal/app"
	"github.com/dziablitsev/shortener/internal/config"
)

var serverAddr string
var shortURLHost string
var serverDebug bool
var shortURLLen int

func main() {
	parseFlags()
	setEnvValues()
	config.SetConfig(buildConfig())

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func buildConfig() (config.ServerConfig, config.ShortURLConfig) {
	var serverConfigBuilder config.ServerConfigBuilder
	serverConfig := serverConfigBuilder.
		WithAddr(serverAddr).
		WithDebug(serverDebug).
		GetConfig()

	var shortURLConfigBuilder config.ShortURLConfigBuilder
	shortURLConfig := shortURLConfigBuilder.
		WithHost(shortURLHost).
		WithLen(shortURLLen).
		GetConfig()

	return serverConfig, shortURLConfig
}
