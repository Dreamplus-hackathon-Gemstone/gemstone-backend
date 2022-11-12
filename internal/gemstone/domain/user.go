package domain

import "time"

type User struct {
	ID            uint        `json:"id" gorm:"primaryKey;autoIncrement;index"`
	Name          string      `json:"name,omitempty" gorm:"not_null;"`
	Nickname      string      `json:"nickname" gorm:"not_null;unique;"`
	Account       string      `json:"account,omitempty" gorm:"not_null;unique;"`
	Email         string      `json:"email,omitempty" gorm:"not_null;unique"`
	Password      string      `json:"password,omitempty" gorm:"not_null;"`
	MyProposals   []Proposal  `gorm:"foreignKey:MakerID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"proposals,omitempty" `
	MyMovies      []Movie     `gorm:"foreignKey:MakerID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"movies,omitempty"`
	MinedProposal []*Proposal `json:"proposal,omitempty" gorm:"many2many:proposal_miners;"`
	MinedMovie    []*Movie    `json:"movie,omitempty" gorm:"many2many:movie_miners;"`
	CreatedAt     time.Time   `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime:nano" json:"updated_at"`
}
