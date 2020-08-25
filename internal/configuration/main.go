package configuration

type Config struct {
	Port int64
}

func GetDefault() *Config {
	return &Config{
		Port: 8080,
	}
}
