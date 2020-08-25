package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port int64
}

func GetDefault() *Config {
	return &Config{
		Port: 8080,
	}
}

func ReadFile(filename string) (*Config, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = json.Unmarshal(f, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
