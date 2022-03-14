package backend

import (
	"context"
	"errors"
	"time"

	"github.com/HotPotatoC/my-unsplash/backend/repository"
)

var (
	ErrSearchPhotosInvalidLimitSize = errors.New("error search photos invalid limit size")
)

const (
	DefaultSearchPhotosLimit = 12
)

func (b backend) SearchPhotos(ctx context.Context, query string, limit int64, createdAtCursor string) ([]repository.Photo, error) {
	if limit < 1 {
		limit = DefaultSearchPhotosLimit
	}

	if createdAtCursor == "" {
		createdAtCursor = time.Now().Format(time.RFC3339)
	}

	return b.photoRepo.SearchPhotos(ctx, query, limit, createdAtCursor)
}
