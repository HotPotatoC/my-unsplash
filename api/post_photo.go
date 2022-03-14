package api

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/HotPotatoC/my-unsplash/backend"
	"github.com/HotPotatoC/my-unsplash/internal/json"
	"github.com/HotPotatoC/my-unsplash/internal/logger"
)

func (h handler) PostPhoto(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyStr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			ReplyJSON(w, http.StatusBadRequest, JSON{
				"message": "Invalid request body",
			})
			return
		}

		var photo backend.PostPhotoParams

		if err := json.Unmarshal(bodyStr, &photo); err != nil {
			ReplyJSON(w, http.StatusBadRequest, JSON{
				"message": "Invalid request body",
			})
			return
		}

		if err := h.backend.PostPhoto(ctx, photo); err != nil {
			logger.S().Error(err)
			ReplyJSON(w, http.StatusUnprocessableEntity, JSON{
				"message": err.Error(),
			})
			return
		}

		ReplyJSON(w, http.StatusCreated, JSON{
			"message": "Photo uploaded successfully",
		})
	})
}
