package config

import "github.com/kelseyhightower/envconfig"

// Config
type Config struct {
	Db     Db
	Server Server
	File   File
}

// Db
type Db struct {
	Uri      string `envconfig:"DB_URI" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

// Server
type Server struct {
	Port int `envconfig:"SERVER_PORT" required:"true"`
}

// File
type File struct {
	SaveDir string `envconfig:"FILE_SAVE_DIR" required:"true"`
}

// New - contructor
func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.Db); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	if err := envconfig.Process("file", &cfg.File); err != nil {
		return nil, err
	}

	return cfg, nil
}
