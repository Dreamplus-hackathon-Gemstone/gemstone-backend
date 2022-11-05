package domain

import "gorm.io/gorm"

type Miner struct {
	gorm.Model `json:"gorm_._model"`
	Account    string      `json:"account,omitempty"`
	Nickname   string      `json:"nickname,omitempty"`
	AuthID     string      `json:"auth_id,omitempty"`
	Proposal   []*Proposal `json:"proposal,omitempty"`
	Movie      []*Movie    `json:"movie,omitempty"`
}
