package db

type Config struct {
	DatabaseURL string `json:"DatabaseURL"`
}

func NewConfig(url string) *Config {
	return &Config{url}
}
