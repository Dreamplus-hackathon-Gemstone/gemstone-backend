package global

import "gemstone-backend/internal/gemstone/domain"

type (
	FindMovieByGenreReq struct {
		GenreID uint `json:"genreId"`
		Offset  int  `json:"offset"`
	}
	FindMovieByGenreRes struct {
		Success bool            `json:"success"`
		Movies  []*domain.Movie `json:"movies"`
		Err     error           `json:"err"`
	}

	FindProposalByGenreReq struct {
		GenreID uint `json:"genreId"`
		Offset  int  `json:"offset"`
	}
	FindProposalByGenreRes struct {
		Success   bool               `json:"success"`
		Proposals []*domain.Proposal `json:"proposals"`
		Err       error              `json:"err"`
	}

	StoreGenreReq struct {
		GenreType string `json:"genre_type"`
	}

	StoreGenreRes struct {
		Success bool  `json:"success"`
		Err     error `json:"err"`
	}
)
