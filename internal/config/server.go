package config

type ServerConfig struct {
	Addr  string
	Debug bool
}

type ServerConfigBuilder struct {
	config ServerConfig
}

func (b ServerConfigBuilder) WithAddr(addr string) ServerConfigBuilder {
	b.config.Addr = addr
	return b
}

func (b ServerConfigBuilder) WithDebug(debug bool) ServerConfigBuilder {
	b.config.Debug = debug
	return b
}

func (b ServerConfigBuilder) GetConfig() ServerConfig {
	return b.config
}
