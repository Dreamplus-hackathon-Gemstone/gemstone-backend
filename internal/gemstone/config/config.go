package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HOST     string
	USER     string
	PASSWORD string
	PORT     string
	NAME     string
}

func NewConfig() *Config {
	HOST := os.Getenv("PSQL_HOST")
	USER := os.Getenv("PSQL_USER")
	PASSWORD := os.Getenv("PSQL_PASSWORD")
	PORT := os.Getenv("PSQL_PORT")
	NAME := os.Getenv("PSQL_NAME")
	return &Config{HOST: HOST, USER: USER, PASSWORD: PASSWORD, PORT: PORT, NAME: NAME}
}

func LoadEnv() error {
	return godotenv.Load(".env")
}
