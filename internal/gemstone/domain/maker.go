package domain

import "gorm.io/gorm"

type Maker struct {
	gorm.Model
	Name        string
	PhoneNumber string
	HomeAddress string
	Password    string
}
