package routers

import (
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	root := chi.NewRouter()

	return root
}
