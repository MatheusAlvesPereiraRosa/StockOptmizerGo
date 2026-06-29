package handler

import (
	"net/http"

	"gear-priority-api/internal/response"
	"gear-priority-api/internal/service"
)

type RestockHandler struct {
	service *service.RestockService
}

func NewRestockHandler(service *service.RestockService) *RestockHandler {
	return &RestockHandler{
		service: service,
	}
}

func (h *RestockHandler) GetPriorities(
	w http.ResponseWriter,
	r *http.Request,
) {

	priorities, err := h.service.GetPriorities(r.Context())

	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
			"",
		)

		return
	}

	response.JSON(
		w,
		http.StatusOK,
		"Restock priorities calculated successfully",
		priorities,
		"",
	)
}
