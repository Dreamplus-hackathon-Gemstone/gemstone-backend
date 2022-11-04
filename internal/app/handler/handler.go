package handler

import "gorm.io/gorm"

// IHandler ...
type IHandler interface {
	GetTx()
	DB() *gorm.DB
}
