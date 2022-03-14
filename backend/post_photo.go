package backend

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (b backend) PostPhoto(ctx context.Context, params PostPhotoParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	hashedPassword, err := params.HashPassword()
	if err != nil {
		return err
	}

	return b.photoRepo.AddPhoto(ctx, params.Label, params.URL, string(hashedPassword))
}

type PostPhotoParams struct {
	Label    string `json:"label"`
	URL      string `json:"url"`
	Password string `json:"password"`
}

func (param PostPhotoParams) Validate() error {
	if param.Label == "" {
		return errors.New("missing label")
	}

	if len(param.Label) > 255 {
		return errors.New("the maximum length of label is 255")
	}

	if param.URL == "" {
		return errors.New("missing url")
	}

	return nil
}

func (param PostPhotoParams) HashPassword() ([]byte, error) {
	digest, err := bcrypt.GenerateFromPassword([]byte(param.Password), 10)
	if err != nil {
		return nil, err
	}

	return digest, nil
}
