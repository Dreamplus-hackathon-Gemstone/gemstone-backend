package domain

import "gorm.io/gorm"

type Maker struct {
	gorm.Model  `json:"gorm_._model"`
	Name        string     `json:"name,omitempty"`
	PhoneNumber string     `json:"phone_number,omitempty"`
	HomeAddress string     `json:"home_address,omitempty"`
	Password    string     `json:"password,omitempty"`
	Proposals   []Proposal `json:"proposals,omitempty" gorm:"foreignKey:Proposer;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
