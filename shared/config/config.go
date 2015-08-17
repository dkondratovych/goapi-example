package config

import (
	"os"
	"errors"
	"encoding/json"
)

var (
	ErrorConfigFileNotFound = errors.New("Unable to find config file")
)

type Config struct {
	Database	*Database 		`json:"database"`
	Application *Application	`json:"application"`
	Server		*Server			`json:"server`
	Facebook 	*Facebook		`json:"facebook"`
	ConfigPath	string			`json:-`
}


func (c *Config) Load(configPath string) error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return ErrorConfigFileNotFound
	}

	f, err := os.Open(configPath)
	defer f.Close()

	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(f)

	if err := jsonParser.Decode(&c); err != nil {
		return err
	}

	c.ConfigPath = configPath

	return nil
}

func NewConfig() *Config {
	return &Config{}
}


