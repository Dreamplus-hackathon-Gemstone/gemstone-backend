package repository

import "gorm.io/gorm"

type Transactor struct {
	client *gorm.DB
}
