package config

import "github.com/kelseyhightower/envconfig"

// Config
type Config struct {
	Server Server
	Store  Store
}

// Server
type Server struct {
	Host string `envconfig:"SERVER_HOST" required:"true"`
}

// Store
type Store struct {
	Url string `envconfig:"STORE_URL" required:"true"`
}

// New - contructor
func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	if err := envconfig.Process("store", &cfg.Store); err != nil {
		return nil, err
	}

	return cfg, nil
}
