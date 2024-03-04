package config

import (
	"github.com/BurntSushi/toml"
	"github.com/creasty/defaults"
)

type (
	Config struct {
		Version uint8 `toml:"version" default:"1"`

		// Logger     struct {
		// 	Debug          bool   `toml:"debug"`
		// 	LoggerEncoding string `toml:"logger_encoding" default:"json"`
		// }
		DSNInfo
	}
)

type DSNInfo struct {
	Driver   string `toml:"driver"`
	Scheme   string `toml:"scheme"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"name"`
	Params   string `toml:"params"`
}

func LoadConfigFromFile(filename string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return nil, err
	}
	if err := defaults.Set(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
