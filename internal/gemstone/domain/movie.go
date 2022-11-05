package domain

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name         string
	Description  string
	ContentID    string
	ThumbnailURI string
	MovieURI     string
}
