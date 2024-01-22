package routers

import (
	"webapp-core/core/domain/health"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func NewProductsRouter(log *zerolog.Logger) chi.Router {
	root := chi.NewRouter()

	root.Get("/", health.Read)

	return root
}
