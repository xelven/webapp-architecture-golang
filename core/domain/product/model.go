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

type Products []*Product

// Request
type PostProductForm struct {
	Name          string `json:"name" form:"required,max=255"`
	Creator       string `json:"creator" form:"required,alpha_space,max=255"`
	PublishedDate string `json:"published_date" form:"required,datetime=2006-01-02"`
	ImageURL      string `json:"image_url" form:"url"`
	Description   string `json:"description"`
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
