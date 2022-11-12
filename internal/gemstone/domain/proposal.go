package domain

import (
	"time"
)

type Proposal struct {
	ID                      uint          `gorm:"primaryKey;autoIncrement;index;"`
	GenreID                 uint          `json:"genre_id,omitempty"`
	MakerID                 uint          `json:"maker_id,omitempty"`
	ContentID               string        `json:"content_id,omitempty"`
	ThumbnailURI            string        `json:"thumbnail_uri,omitempty"`
	DocumentURI             string        `json:"document_uri,omitempty"`
	Title                   string        `json:"title,omitempty"`
	Description             string        `json:"description,omitempty"`
	TargetPrice             uint64        `json:"target_price,omitempty"`
	Deadline                time.Time     `json:"deadline"`
	ExpectedReleaseYear     time.Time     `json:"expected_release_year"`
	EstimatedProductionTime time.Duration `json:"estimated_production_time,omitempty"`
	Miners                  []*User       `json:"miners" gorm:"many2many:proposal_users;"`
	CreatedAt               time.Time     `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt               time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProposalUser struct {
	ProposalID uint `json:"proposal_id,omitempty"`
	UserID     uint `json:"user_id,omitempty"`
}
