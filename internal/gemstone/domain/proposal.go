package domain

import (
	"gorm.io/gorm"
	"time"
)

type Proposal struct {
	gorm.Model              `json:"gorm_._model"`
	ContentID               string        `json:"content_id,omitempty"`
	ThumbnailURI            string        `json:"thumbnail_uri,omitempty"`
	DocumentURI             string        `json:"document_uri,omitempty"`
	GenreType               string        `json:"genre_type,omitempty"`
	Title                   string        `json:"title,omitempty"`
	Proposer                string        `json:"proposer,omitempty"`
	Description             string        `json:"description,omitempty"`
	TargetPrice             uint64        `json:"target_price,omitempty"`
	Deadline                time.Time     `json:"deadline"`
	ExpectedReleaseYear     time.Time     `json:"expected_release_year"`
	EstimatedProductionTime time.Duration `json:"estimated_production_time,omitempty"`
	Miners                  []*Miner      `json:"miners" gorm:"foreignKey:ProposalID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
