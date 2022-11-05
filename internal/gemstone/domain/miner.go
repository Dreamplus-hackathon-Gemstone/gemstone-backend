package domain

import "gorm.io/gorm"

type Miner struct {
	gorm.Model
	account string
	AuthID  string
}
