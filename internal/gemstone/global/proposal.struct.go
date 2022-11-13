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
		Success   bool               `json:"success,omitempty"`
		Proposals []*domain.Proposal `json:"proposals,omitempty"`
	}

	FindProposalReq struct {
		CID string `json:"cid,omitempty"`
	}

	FindProposalRes struct {
		Success  bool             `json:"success,omitempty"`
		Proposal *domain.Proposal `json:"proposal,omitempty"`
		Err      error            `json:"err"`
	}

	StoreProposalReq struct {
		GenreID                 uint      `json:"genre_id,omitempty"`
		MakerID                 uint      `json:"maker_id,omitempty"`
		ContentID               string    `json:"content_id,omitempty"`
		ThumbnailURI            string    `json:"thumbnail_uri,omitempty"`
		DocumentURI             string    `json:"document_uri,omitempty"`
		Title                   string    `json:"title,omitempty"`
		Description             string    `json:"description,omitempty"`
		TargetPrice             uint64    `json:"target_price,omitempty"`
		Deadline                time.Time `json:"deadline"`
		ExpectedReleaseYear     time.Time `json:"expected_release_year"`
		EstimatedProductionTime string    `json:"estimated_production_time,omitempty"`
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
