package config

var Server ServerConfig

var ShortURL ShortURLConfig

func SetConfig(serverConfig ServerConfig, shortURLConfig ShortURLConfig) {
	Server = serverConfig
	ShortURL = shortURLConfig
}
