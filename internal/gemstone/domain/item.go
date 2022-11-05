package domain

import "gorm.io/gorm"

type Item struct {
	gorm.Model `json:"gorm_._model"`
	Title      string `json:"title,omitempty"`
	// Genre
	Description string `json:"description,omitempty"`
}
