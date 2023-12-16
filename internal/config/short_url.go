package config

type ShortURLConfig struct {
	Host string
	Len  int
}

type ShortURLConfigBuilder struct {
	config ShortURLConfig
}

func (b ShortURLConfigBuilder) WithHost(host string) ShortURLConfigBuilder {
	b.config.Host = host
	return b
}

func (b ShortURLConfigBuilder) WithLen(len int) ShortURLConfigBuilder {
	b.config.Len = len
	return b
}

func (b ShortURLConfigBuilder) GetConfig() ShortURLConfig {
	return b.config
}
