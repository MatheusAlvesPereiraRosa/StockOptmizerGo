package handler

import (
	"encoding/json"
	"net/http"

	"gear-priority-api/internal/domain"
	"gear-priority-api/internal/response"
	"gear-priority-api/internal/service"

	"github.com/go-chi/chi/v5"

	"github.com/google/uuid"
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
		response.Error(w, http.StatusInternalServerError, "Failed to create gear", err.Error())
		return
	}

	response.JSON(
		w,
		http.StatusCreated,
		"Gear created successfully",
		gear,
		"",
	)
}

func (h *GearHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	w.Header().Set("Context-Type", "application/json")

	print(category)

	if category != "" {
		gears, err := h.service.FindByCategory(r.Context(), category)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(gears)
	} else {
		gears, err := h.service.FindAll(r.Context())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(gears)
	}
}

func (h *GearHandler) FindByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	gear, err := h.service.FindByID(r.Context(), idParam)

	if err != nil {
		response.Error(w, http.StatusNotFound, "Gear not found", err.Error())

		return
	}

	response.JSON(w, http.StatusOK, "Gear found", gear, "")
}

func (h *GearHandler) Update(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid UUID", err.Error())

		return
	}

	var gear domain.Gear

	if err := json.NewDecoder(r.Body).Decode(&gear); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body", err.Error())

		return
	}

	gear.ID = id

	err = h.service.Update(r.Context(), &gear)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update gear", err.Error())

		return
	}

	response.JSON(
		w,
		http.StatusOK,
		"Gear updated successfully",
		gear,
		"",
	)
}

func (h *GearHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.service.Delete(r.Context(), id)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to delete gear", err.Error())
		return
	}

	response.JSON(
		w,
		http.StatusOK,
		"Gear deleted successfully",
		nil,
		"",
	)
}
