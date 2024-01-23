package product

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DB struct
type Product struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Name          string    `gorm:"type:text"`
	Creator       string    `gorm:"type:text"`
	PublishedDate time.Time
	ImageURL      string `gorm:"type:text"`
	Description   string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

// Request
type PostProductForm struct {
	Name          string `json:"name" form:"required,max=255"`
	Creator       string `json:"creator" form:"required,alpha_space,max=255"`
	PublishedDate string `json:"published_date" form:"required,datetime=2006-01-02"`
	ImageURL      string `json:"image_url" form:"url"`
	Description   string `json:"description"`
}

// Response struct
type ProductResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Creator       string    `json:"creator"`
	PublishedDate string    `json:"published_date"`
	ImageURL      string    `json:"image_url"`
	Description   string    `json:"description"`
	UpdatedAt     time.Time `json:"last_update"`
}

// request body to db model
func (f *PostProductForm) FormToModel() *Product {
	pubDate, _ := time.Parse("2016-12-23", f.PublishedDate)

	return &Product{
		Name:          f.Name,
		Creator:       f.Creator,
		PublishedDate: pubDate,
		ImageURL:      f.ImageURL,
		Description:   f.Description,
	}
}

type Products []*Product

// db model to response body
func (b *Product) ProductToResponse() *ProductResponse {
	return &ProductResponse{
		ID:            b.ID.String(),
		Name:          b.Name,
		Creator:       b.Creator,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageURL:      b.ImageURL,
		Description:   b.Description,
		UpdatedAt:     b.UpdatedAt,
	}
}

func (products Products) ProductsToResponse() []*ProductResponse {
	productsResponse := make([]*ProductResponse, len(products))
	for i, v := range products {
		productsResponse[i] = v.ProductToResponse()
	}

	return productsResponse
}
