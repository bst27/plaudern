package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port           int64
	AllowedOrigins []string
	DatabaseFile   string
}

func GetDefault() *Config {
	origins := []string{
		"http://example.com",
		"https://example.com/",
	}

	return &Config{
		Port:           8080,
		AllowedOrigins: origins,
		DatabaseFile:   "plaudern-data.db",
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
