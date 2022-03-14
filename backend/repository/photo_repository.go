package repository

import (
	"context"
	"time"

	"github.com/HotPotatoC/my-unsplash/clients"
)

type Photo struct {
	ID        int64     `json:"id"`
	Label     string    `json:"label"`
	URL       string    `json:"photo_url"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoRepository struct {
	db *clients.PostgresClient
}

func NewPhotoRepository(db *clients.PostgresClient) PhotoRepository {
	return PhotoRepository{db: db}
}

const addPhotoQuery = `
INSERT INTO photos (
	label,
	url,
	password,
	created_at
) VALUES (
	$1,
	$2,
	$3,
	$4
)`

func (r *PhotoRepository) AddPhoto(ctx context.Context, label, url, password string) error {
	_, err := r.db.Exec(ctx, addPhotoQuery, label, url, password, time.Now())
	if err != nil {
		return err
	}

	return nil
}

const listPhotosQuery = `
SELECT
	id,
	label,
	url,
	password,
	created_at
FROM photos
WHERE created_at < $1
ORDER BY created_at DESC
LIMIT $2`

func (r *PhotoRepository) ListPhotos(ctx context.Context, limit int64, createdAtCursor string) ([]Photo, error) {
	var photos []Photo

	rows, err := r.db.Query(ctx, listPhotosQuery, createdAtCursor, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		err := rows.Scan(&photo.ID, &photo.Label, &photo.URL, &photo.Password, &photo.CreatedAt)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return photos, nil
}

const findPhotoQuery = `
SELECT
	id,
	label,
	url,
	password,
	created_at
FROM photos
WHERE id = $1`

func (r *PhotoRepository) FindPhoto(ctx context.Context, id int64) (Photo, error) {
	var photo Photo

	err := r.db.QueryRow(ctx, findPhotoQuery, id).Scan(&photo.ID, &photo.Label, &photo.URL, &photo.Password, &photo.CreatedAt)
	if err != nil {
		return Photo{}, err
	}

	return photo, nil
}

const deletePhotoQuery = `DELETE FROM photos WHERE id = $1`

func (r *PhotoRepository) DeletePhoto(ctx context.Context, id int64) (int64, error) {
	n, err := r.db.Exec(ctx, deletePhotoQuery, id)
	if err != nil {
		return 0, err
	}

	return n, nil
}

const searchPhotoQuery = `
SELECT
	id,
	label,
	label_tsvector,
	url,
	password,
	created_at
FROM photos
WHERE label_tsvector @@ plainto_tsquery($1) AND created_at < $2
ORDER BY label_tsvector DESC
LIMIT $3`

func (r *PhotoRepository) SearchPhotos(ctx context.Context, query string, limit int64, createdAtCursor string) ([]Photo, error) {
	var photos []Photo

	rows, err := r.db.Query(ctx, searchPhotoQuery, query, createdAtCursor, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo Photo
		err := rows.Scan(&photo.ID, &photo.Label, nil, &photo.URL, &photo.Password, &photo.CreatedAt)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return photos, nil
}
