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

func (r *Querier) getList() (Products, error) {
	products := make([]*Product, 0)
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
