package config

var Server ServerConfig
var ShortURL ShortURLConfig

var serverAddr string
var shortURLHost string
var serverDebug bool
var shortURLLen int

func SetConfig(serverConfig ServerConfig, shortURLConfig ShortURLConfig) {
	Server = serverConfig
	ShortURL = shortURLConfig
}

func BuildConfig() (ServerConfig, ShortURLConfig) {
	ParseFlags()
	SetEnvValues()

	var serverConfigBuilder ServerConfigBuilder
	serverConfig := serverConfigBuilder.
		WithAddr(serverAddr).
		WithDebug(serverDebug).
		GetConfig()

	var shortURLConfigBuilder ShortURLConfigBuilder
	shortURLConfig := shortURLConfigBuilder.
		WithHost(shortURLHost).
		WithLen(shortURLLen).
		GetConfig()

	return serverConfig, shortURLConfig
}
