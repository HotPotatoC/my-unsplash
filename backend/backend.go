package backend

import (
	"context"

	"github.com/HotPotatoC/my-unsplash/backend/repository"
	"github.com/HotPotatoC/my-unsplash/clients"
)

type Backend interface {
	PostPhoto(ctx context.Context, params PostPhotoParams) error
	ListPhotos(ctx context.Context, limit int64, createdAtCursor string) ([]repository.Photo, error)
	DeletePhoto(ctx context.Context, id int64, password string) (int64, error)
	SearchPhotos(ctx context.Context, query string, limit int64, createdAtCursor string) ([]repository.Photo, error)
}

type backend struct {
	photoRepo repository.PhotoRepository
}

func New(clients clients.Clients) Backend {
	return backend{
		photoRepo: repository.NewPhotoRepository(clients.PostgresClient),
	}
}
