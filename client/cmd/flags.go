package main

import (
	"flag"
)

type Config struct {
	Endpoint  string
	CryptoKey string
}

func mustReadConfig() Config {
	var config Config

	flag.StringVar(&config.Endpoint, "e", "localhost:8080", "address and port of server to connect to")
	flag.StringVar(&config.CryptoKey, "k", "", "crypto key to store your data securely")
	flag.Parse()

	return config
}
