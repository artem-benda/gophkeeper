package main

import (
	"flag"

	"github.com/caarlos0/env/v10"
)

// Config - конфигурация приложения
type Config struct {
	Endpoint    string `env:"RUN_ADDRESS"`
	DatabaseDSN string `env:"DATABASE_URI"`
	Salt        string `env:"SALT"`
}

// mustReadConfig - прочитать конфигурацию, переданную при запуске приложения
func mustReadConfig() Config {
	var config Config

	flag.StringVar(&config.Endpoint, "e", "localhost:8080", "address and port of server")
	flag.StringVar(&config.DatabaseDSN, "d", "postgres://gophkeeper:gophkeeper123@localhost:5432/gophkeeper?sslmode=disable", "Database connection URL in pgx format, for ex. postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10")
	flag.StringVar(&config.Salt, "s", "BPjkLEqJfARvsYGW++WRcnCjxHyZsrnxXd/qdzpWIaE=", "salt in base64std format, using for hashing passwords, at least 8 bytes is recommended by the RFC")

	flag.Parse()
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
