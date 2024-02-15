package config

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: ":8080",
		DatabaseURL:   "postgres://postgres:password1@localhost:5433/postgres?sslmode=disable",
	}
}
