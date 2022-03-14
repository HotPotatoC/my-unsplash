package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/HotPotatoC/my-unsplash/internal/logger"
)

func (h handler) ListPhotos(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")

		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil && limitStr != "" {
			ReplyJSON(w, http.StatusBadRequest, JSON{
				"message": "Invalid limit provided",
			})
			return
		}

		createdAtCursor := r.URL.Query().Get("cursor")

		photos, err := h.backend.ListPhotos(ctx, limit, createdAtCursor)
		if err != nil {
			logger.S().Error(err)
			ReplyJSON(w, http.StatusInternalServerError, JSON{
				"message": "There was a problem on our side",
			})
			return
		}

		ReplyJSON(w, http.StatusOK, JSON{
			"items":       photos,
			"total_items": len(photos),
		})
	})
}
