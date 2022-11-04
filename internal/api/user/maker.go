package user

import "gorm.io/gorm"

// Maker ...
type Maker struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Nickname    string `gorm:"nickname" json:"nickname"`
	Email       string `gorm:"email" json:"email"`
	PhoneNumber string `gorm:"phone_number" json:"phone_number"`
}

func (maker *Maker) TableName() string {
	return "maker"
}
