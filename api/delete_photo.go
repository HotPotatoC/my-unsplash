package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/HotPotatoC/my-unsplash/backend"
	"github.com/HotPotatoC/my-unsplash/internal/logger"
	"github.com/go-chi/chi/v5"
)

func (h handler) DeletePhoto(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			ReplyJSON(w, http.StatusBadRequest, JSON{
				"message": "Invalid id provided",
			})
			return
		}

		password := r.URL.Query().Get("password")

		deletedN, err := h.backend.DeletePhoto(ctx, id, password)
		if err != nil {
			switch err {
			case backend.ErrPhotoDoesNotExist:
				ReplyJSON(w, http.StatusNotFound, JSON{
					"message": "Photo does not exist",
				})
				return
			case backend.ErrForbiddenPhotoDeletionAction:
				ReplyJSON(w, http.StatusForbidden, JSON{
					"message": "Forbidden photo deletion action",
				})
				return
			default:
				logger.S().Error(err)
				ReplyJSON(w, http.StatusInternalServerError, JSON{
					"message": "There was a problem on our side",
				})
				return
			}
		}

		ReplyJSON(w, http.StatusOK, JSON{
			"deleted": deletedN,
			"message": "Photo deleted",
		})
	})
}
