package handler

import (
	"net/http"

	"gear-priority-api/internal/service"
	"gear-priority-api/internal/utils"
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
		utils.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
			"",
		)

		return
	}

	utils.JSON(
		w,
		http.StatusOK,
		"Restock priorities calculated successfully",
		priorities,
		"",
	)
}
