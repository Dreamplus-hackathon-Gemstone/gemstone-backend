package domain

import "time"

type Movies struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;index;"`
	MakerID      uint      `json:"maker_id,omitempty" gorm:"not_null"`
	GenreID      uint      `json:"genre_id,omitempty" gorm:"not_null"`
	ContentID    string    `json:"content_id,omitempty" gorm:"not_null;index;unique"`
	Description  string    `json:"description,omitempty" gorm:"not_null"`
	ThumbnailURI string    `json:"thumbnail_uri,omitempty"`
	MovieURI     string    `json:"movie_uri,omitempty" `
	Miners       []*Miners `gorm:"many2many:movie_miners;" json:"miners"`
	CreatedAt    time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MovieMiners struct {
	MoviesID uint `json:"movies_id,omitempty"`
	MinersID uint `json:"miner_id,omitempty"`
}
