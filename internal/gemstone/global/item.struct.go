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

	FindManyReq struct {
		Offset uint `json:"offset"`
	}

	FindManyMovieRes struct {
		Success bool `json:"success"`
		Movie   []*domain.Movie
	}

	FindManyProposalRes struct {
		Success bool `json:"success"`
		Movie   []*domain.Proposal
	}
)
