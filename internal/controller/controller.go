package router

import (
	"gear-priority-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *handler.GearHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/gears", func(r chi.Router) {
		r.Get("/", handler.FindAll)

		r.Post("/", handler.Create)
	})

	return r
}
