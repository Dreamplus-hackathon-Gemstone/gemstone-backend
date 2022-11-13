package global

import "gemstone-backend/internal/gemstone/domain"

type (
	FindMovieReq struct {
		CID string `json:"cid"`
	}

	FindMovieRes struct {
		Success bool          `json:"success"`
		Movie   *domain.Movie `json:"movie"`
		Err     error         `json:"err"`
	}

	FindManyMovieReq struct {
		Offset  int  `json:"offset"`
		GenreID uint `json:"genre_id"`
	}

	FindManyMovieRes struct {
		Success bool            `json:"success"`
		Movies  []*domain.Movie `json:"movies"`
		Err     error           `json:"err"`
	}

	StoreMovieReq struct {
		MakerID      uint   `json:"maker_id"`
		GenreID      uint   `json:"genre_id"`
		ContentID    string `json:"content_id"`
		Description  string `json:"description"`
		ThumbnailURI string `json:"thumbnail_uri"`
		MovieURI     string `json:"movie_uri"`
	}

	StoreMovieRes struct {
		Success bool  `json:"success"`
		Err     error `json:"err"`
	}
)
