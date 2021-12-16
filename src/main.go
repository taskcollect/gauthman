package main

import (
	"errors"
	"log"
	"main/auth"
	"main/handlers"
	"net/http"
	"os"
)

type ServerConfig struct {
	BindAddr string
	Secrets  *auth.OAuth2Secrets
}

func getSecretsFromEnv() (*auth.OAuth2Secrets, error) {
	id := os.Getenv("CLIENT_ID")
	if id == "" {
		return nil, errors.New("CLIENT_ID not set or empty in environment")
	}

	secret := os.Getenv("CLIENT_SECRET")
	if secret == "" {
		return nil, errors.New("CLIENT_SECRET not set or empty in environment")
	}

	return &auth.OAuth2Secrets{
		ClientID:     id,
		ClientSecret: secret,
	}, nil
}

// server config, values here will get overriden by env
var config = ServerConfig{
	BindAddr: "0.0.0.0:2000",
	Secrets:  nil,
}

func makeMux(secrets *auth.OAuth2Secrets) *http.ServeMux {
	mux := http.NewServeMux()

	handler := handlers.NewBaseHandler(secrets)

	mux.HandleFunc("/v1/exchange", handler.ExchangeToken)

	return mux
}

func configure(c *ServerConfig) {
	bindAddr, exists := os.LookupEnv("BIND_ADDR")
	if exists {
		if bindAddr == "" {
			log.Fatalln("(cfg) empty bind address supplied, cannot bind")
		}
		c.BindAddr = bindAddr
	} else {
		log.Printf("(cfg) no bind address supplied, defaulting to '%s'", c.BindAddr)
	}

	secrets, err := getSecretsFromEnv()
	if err != nil {
		log.Fatalln("(cfg) error in credential init:", err.Error())
	}

	c.Secrets = secrets
}

func main() {
	log.Printf("Initializing config from environment variables...")

	configure(&config)

	log.Printf("Starting server binded to %s...", config.BindAddr)

	mux := makeMux(config.Secrets)
	http.ListenAndServe(config.BindAddr, handlers.RequestLogger(mux))

	log.Println("Server exited. Cleaning up...")
}
