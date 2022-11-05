package domain

import "time"

type Token struct {
	TokenID   string    `json:"token_id,omitempty" gorm:"primaryKey;autoIncrement:false"`
	TokenURI  string    `json:"token_uri,omitempty"`
	CreatedAt time.Time `gorm:"not_null"`
	UpdatedAt time.Time `gorm:"not_null"`
}
