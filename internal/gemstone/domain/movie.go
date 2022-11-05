package domain

import "time"

type Movie struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;index;"`
	Nickname     string    `json:"name,omitempty" gorm:"not_null"`
	Description  string    `json:"description,omitempty" gorm:"not_null"`
	ContentID    string    `json:"content_id,omitempty" gorm:"not_null;index;unique"`
	ThumbnailURI string    `json:"thumbnail_uri,omitempty" gorm:"not_null"`
	MovieURI     string    `json:"movie_uri,omitempty" gorm:"not_null"`
	Miners       []*Miners `gorm:"many2many:movie_miners;" json:"miners"`
	CreatedAt    time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MovieMiners struct {
	MoviesID uint `json:"movies_id,omitempty"`
	MinersID uint `json:"miner_id,omitempty"`
}
