package main

import (
	"context"
	"os"

	"github.com/HotPotatoC/my-unsplash/api"
	"github.com/HotPotatoC/my-unsplash/backend"
	"github.com/HotPotatoC/my-unsplash/clients"
	"github.com/HotPotatoC/my-unsplash/internal/logger"
)

func main() {
	logger.Init(os.Getenv("APP_ENV") != "production")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clients, err := clients.Init(ctx, &clients.Options{
		PostgreSQLConnString: os.Getenv("POSTGRESQL_URL"),
	})
	if err != nil {
		logger.S().Error(err)
		return
	}

	backend := backend.New(clients)

	api.NewServer(ctx, backend).Serve()
}
