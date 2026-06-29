package controller

import (
	"gear-priority-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(
	gearHandler *handler.GearHandler,
	restockHandler *handler.RestockHandler,
) *chi.Mux {

	r := chi.NewRouter()

	GearRouter(r, gearHandler)
	RestockRouter(r, restockHandler)

	return r
}
