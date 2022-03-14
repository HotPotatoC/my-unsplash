package backend

import (
	"context"
	"errors"
	"time"

	"github.com/HotPotatoC/my-unsplash/backend/repository"
)

var (
	ErrListPhotosInvalidLimitSize = errors.New("error list photos invalid limit size")
)

const (
	DefaultListPhotosLimit = 12
)

func (b backend) ListPhotos(ctx context.Context, limit int64, createdAtCursor string) ([]repository.Photo, error) {
	if limit < 1 {
		limit = DefaultListPhotosLimit
	}

	if createdAtCursor == "" {
		createdAtCursor = time.Now().Format(time.RFC3339)
	}

	return b.photoRepo.ListPhotos(ctx, limit, createdAtCursor)
}
