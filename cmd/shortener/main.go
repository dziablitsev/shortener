package main

import (
	"github.com/dziablitsev/shortener/internal/app"
	"github.com/dziablitsev/shortener/internal/config"
)

func main() {
	parseFlags()
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
