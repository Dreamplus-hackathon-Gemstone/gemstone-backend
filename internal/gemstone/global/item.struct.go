package global

import "gemstone-backend/internal/gemstone/domain"

type (
	FindItemReq struct {
		ContentID string `json:"content_id"`
	}

	FindMovieRes struct {
		Success bool `json:"success"`
		Movie   *domain.Movie
	}

	FindManyMovieRes struct {
		Success bool `json:"success"`
		Movie   []*domain.Movie
	}
)
