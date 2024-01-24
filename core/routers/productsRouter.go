package routers

import (
	"webapp-core/core/domain/product"
	"webapp-core/core/routers/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func NewProductsRouter(log *zerolog.Logger, db *gorm.DB) chi.Router {
	products := chi.NewRouter()

	products.Use(middleware.ContentTypeJSON)

	productAPI := product.New(log, db)

	products.Get("/", productAPI.GetList)
	products.Post("/", productAPI.Create)

	return products
}
