package main

import (
	"flag"
)

type Config struct {
	CryptoKey        string
	DatabaseFilePath string
	LogLevel         string
}

func mustReadConfig() Config {
	var config Config

	flag.StringVar(&config.CryptoKey, "key", "", "crypto key to store your data securely")
	flag.StringVar(&config.DatabaseFilePath, "db", "./default.db", "path or connection url to local sqlite database")
	flag.StringVar(&config.LogLevel, "l", "debug", "logging level: debug, info, warn, error, dpanic, panic, fatal")
	flag.Parse()

	return config
}
