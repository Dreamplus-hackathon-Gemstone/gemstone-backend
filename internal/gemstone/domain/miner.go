package domain

import "time"

type Miners struct {
	ID        uint         `gorm:"primaryKey;autoIncrement;index;"`
	Account   string       `json:"account,omitempty" gorm:"not_null"`
	Nickname  string       `json:"nickname,omitempty" gorm:"not_null"`
	AuthID    string       `json:"auth_id,omitempty" gorm:"not_null"`
	Proposal  []*Proposals `json:"proposal,omitempty" gorm:"many2many:proposal_miners;"`
	Movie     []*Movie     `json:"movie,omitempty" gorm:"many2many:movie_miners;"`
	CreatedAt time.Time    `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
