package routers

import (
	"webapp-core/core/domain/health"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func New(log *zerolog.Logger, db *gorm.DB) *chi.Mux {
	root := chi.NewRouter()

	root.Get("/health", health.Read)

	root.Mount("/products", NewProductsRouter(log, db))

	return root
}
