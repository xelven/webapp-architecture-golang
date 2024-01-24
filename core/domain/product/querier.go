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

func (query *Querier) GetList() (Products, error) {
	products := make([]*Product, 0)
	/*
		// orm query
		if err := query.db.Find(&products).Error; err != nil {
			return nil, err
		}
	*/
	// https://gorm.io/docs/sql_builder.html
	/*
		type Result struct {
			ID   int
			Name string
			Age  int
		}

		var result Result
		db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)

	*/
	if err := query.db.Raw("SELECT * FROM products").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (query *Querier) List(offset int, limit int) (Products, error) {
	products := make([]*Product, 0)
	if err := query.db.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
