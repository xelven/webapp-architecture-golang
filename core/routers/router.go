package routers

import (
	"webapp-core/core/domain/health"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func New(log *zerolog.Logger) *chi.Mux {
	root := chi.NewRouter()

	root.Get("/health", health.Read)

	return root
}
