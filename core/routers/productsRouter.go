package routers

import (
	"webapp-core/core/domain/product"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func NewProductsRouter(log *zerolog.Logger, db *gorm.DB) chi.Router {
	products := chi.NewRouter()

	productAPI := product.New(log, db)

	products.Get("/", productAPI.GetList)
	products.Post("/", productAPI.Create)

	return products
}
