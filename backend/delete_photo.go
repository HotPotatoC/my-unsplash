package backend

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPhotoDoesNotExist            = errors.New("error photo does not exist")
	ErrForbiddenPhotoDeletionAction = errors.New("error forbidden photo deletion action")
)

func (b backend) DeletePhoto(ctx context.Context, id int64, password string) (int64, error) {
	photo, err := b.photoRepo.FindPhoto(ctx, id)
	if err != nil {
		return 0, ErrPhotoDoesNotExist
	}

	err = bcrypt.CompareHashAndPassword([]byte(photo.Password), []byte(password))

	if err != nil && photo.Password != "" {
		return 0, ErrForbiddenPhotoDeletionAction
	}

	return b.photoRepo.DeletePhoto(ctx, id)
}
