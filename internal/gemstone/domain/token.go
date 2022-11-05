package domain

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	tokenID  string
	tokenURI string
}
