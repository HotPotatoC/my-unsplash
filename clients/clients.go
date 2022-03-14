package clients

import (
	"context"

	"github.com/HotPotatoC/my-unsplash/internal/logger"
)

type Clients struct {
	PostgresClient *PostgresClient
}

func Init(ctx context.Context, opts *Options) (Clients, error) {
	logger.S().Info("Initializing clients")

	opts.Init()

	logger.S().Info("Initializing postgres client...")
	postgresClient, err := NewPostgreSQLClient(ctx, opts.PostgreSQLConnString)
	if err != nil {
		return Clients{}, err
	}

	logger.S().Info("Clients initialized")
	return Clients{
		PostgresClient: postgresClient,
	}, nil
}
