package product

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repository *Repository) Create(product *Product) (*Product, error) {
	if err := repository.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
