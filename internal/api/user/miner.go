package user

import (
	"gorm.io/gorm"
)

// Miner ...
type Miner struct {
	gorm.Model
	Nickname  string `gorm:"nickname" json:"nickname"`
	Account   string `gorm:"account" json:"account"`
	RequestID string `gorm:"request_id" json:"request_id"`
}

func (miner *Miner) TableName() string {
	return "miner"
}
