package domain

import "time"

type Makers struct {
	ID          uint        `gorm:"primaryKey;autoIncrement;index;"`
	Name        string      `json:"name,omitempty" gorm:"not_null;"`
	Nickname    string      `json:"nickname" gorm:"not_null;unique;"`
	PhoneNumber string      `json:"phone_number,omitempty" gorm:"not_null;"`
	HomeAddress string      `json:"home_address,omitempty" gorm:"not_null;"`
	Password    string      `json:"password,omitempty" gorm:"not_null;"`
	Proposals   []Proposals `gorm:"foreignKey:Nickname;references:Nickname;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"proposals,omitempty" `
	Movies      []Movie     `gorm:"foreignKey:Nickname;references:Nickname;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"movies,omitempty"`
	CreatedAt   time.Time   `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
