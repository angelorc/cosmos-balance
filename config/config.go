package config

import (
	"fmt"
	"github.com/naoina/toml"
	"io/ioutil"
)

var DefaultConfigPath = "./config.toml"

type Config struct {
	Bitsong *BitsongConfig `toml:"bitsong"`
	Server  *ServerConfig  `toml:"server"`
}

type BitsongConfig struct {
	GRPC  string `toml:"grpc"`
	Denom string `toml:"denom"`
}

type ServerConfig struct {
	Address string `toml:"address"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		return nil, fmt.Errorf("config file not found")
	}

	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	var cfg Config

	err = toml.Unmarshal(configFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config: %s", err)
	}

	return &cfg, nil
}
