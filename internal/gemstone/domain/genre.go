package domain

import "gorm.io/gorm"

type Genre struct {
	gorm.Model `json:"gorm_._model"`
	GenreType  string `json:"name,omitempty" gorm:"foreignKey:GenreType;references:GenreType;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
