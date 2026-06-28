package handler

import (
	"encoding/json"
	"net/http"

	"gear-priority-api/internal/domain"
	"gear-priority-api/internal/service"
)

type GearHandler struct {
	service *service.GearService
}

func NewGearHandler(service *service.GearService) *GearHandler {
	return &GearHandler{
		service: service,
	}
}

func (h *GearHandler) Create(w http.ResponseWriter, r *http.Request) {
	var gear domain.Gear

	if err := json.NewDecoder(r.Body).Decode(&gear); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	if err := h.service.Create(r.Context(), &gear); err != nil {
		domain.Error(w, http.StatusInternalServerError, "Failed to create gear", err.Error())
		return
	}

	domain.JSON(
		w,
		http.StatusCreated,
		"Gear created successfully",
		gear,
		"",
	)
}

func (h *GearHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	gears, err := h.service.FindAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application/json")

	json.NewEncoder(w).Encode(gears)
}
