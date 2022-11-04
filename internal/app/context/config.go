package context

import "os"

// Config ...
type Config struct {
	Host          string
	Password      string
	Port          string
	User          string
	Name          string
	ServerAddress string
}

// CreateConfigFromEnv ...
func CreateConfigFromEnv() Config {
	var config = Config{}
	config.Host = os.Getenv("PSQL_HOST")
	config.User = os.Getenv("PSQL_USER")
	config.Password = os.Getenv("PSQL_PASSWORD")
	config.Port = os.Getenv("PSQL_PORT")
	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.Name = os.Getenv("PSQO_NAME")
	return config
}
