package global

import (
	"gemstone-backend/internal/gemstone/domain"
	"time"
)

type (
	FindManyProposalReq struct {
		Offset int `json:"offset"`
	}
	FindManyProposalRes struct {
		Success   bool               `json:"success"`
		Proposals []*domain.Proposal `json:"proposals"`
	}

	FindProposalReq struct {
		CID string `json:"cid"`
	}

	FindProposalRes struct {
		Success  bool             `json:"success"`
		Proposal *domain.Proposal `json:"proposal"`
		Err      error            `json:"err"`
	}

	StoreProposalReq struct {
		GenreID                 uint      `json:"genre_id"`
		MakerID                 uint      `json:"maker_id"`
		ContentID               string    `json:"content_id"`
		ThumbnailURI            string    `json:"thumbnail_uri"`
		DocumentURI             string    `json:"document_uri"`
		Title                   string    `json:"title"`
		Description             string    `json:"description"`
		TargetPrice             uint64    `json:"target_price"`
		Deadline                time.Time `json:"deadline"`
		ExpectedReleaseYear     time.Time `json:"expected_release_year"`
		EstimatedProductionTime string    `json:"estimated_production_time"`
	}

	StoreProposalRes struct {
		Success bool
		Err     error
	}

	UpdateProposalReq struct{}
	UpdateProposalRes struct{}

	DeleteProposalReq struct{}
	DeleteProposalRes struct{}
)
