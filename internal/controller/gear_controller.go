package controller

import (
	"gear-priority-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func GearRouter(r chi.Router, handler *handler.GearHandler) {
	r.Route("/gears", func(r chi.Router) {
		r.Get("/", handler.FindAll)

		r.Get("/{id}", handler.FindByID)

		r.Get("/page", handler.FindPaginated)

		r.Post("/", handler.Create)

		r.Put("/{id}", handler.Update)

		r.Delete("/{id}", handler.Delete)
	})
}
