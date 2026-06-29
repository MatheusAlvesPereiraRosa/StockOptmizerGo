package controller

import (
	"gear-priority-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func RestockRouter(r chi.Router, handler *handler.RestockHandler) {
	r.Route("/restock", func(r chi.Router) {
		r.Get("/priorities", handler.GetPriorities)
	})
}
