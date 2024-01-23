package product

import "gorm.io/gorm"

type Querier struct {
	db *gorm.DB
}

func NewQuerier(db *gorm.DB) *Querier {
	return &Querier{
		db: db,
	}
}
