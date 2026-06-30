package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gear-priority-api/internal/dto"
	"gear-priority-api/internal/service"
	"gear-priority-api/internal/utils"

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

func getPaginationParams(r *http.Request) (int, int) {
	page := 1
	limit := 20

	if value := r.URL.Query().Get("page"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			page = parsed
		}
	}

	if value := r.URL.Query().Get("limit"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			limit = parsed
		}
	}

	return page, limit
}

func (h *GearHandler) Create(w http.ResponseWriter, r *http.Request) {
	var gear dto.Gear

	if err := json.NewDecoder(r.Body).Decode(&gear); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	if err := h.service.Create(r.Context(), &gear); err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to create gear")
		return
	}

	utils.JSON(
		w,
		http.StatusCreated,
		"Gear created successfully",
		gear,
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

func (h *GearHandler) FindPaginated(
	w http.ResponseWriter,
	r *http.Request,
) {
	page, limit := getPaginationParams(r)
	category := r.URL.Query().Get("category")

	var (
		result *dto.PaginatedGearsResponse
		err    error
	)

	if category != "" {
		result, err = h.service.FindByCategoryPaginated(
			r.Context(),
			category,
			page,
			limit,
		)
	} else {
		result, err = h.service.FindPaginated(
			r.Context(),
			page,
			limit,
		)
	}

	if err != nil {
		utils.Error(
			w,
			http.StatusInternalServerError,
			"Failed to retrieve paginated gears",
		)
		return
	}

	utils.Pagination(
		w,
		http.StatusOK,
		result,
	)
}

func (h *GearHandler) FindByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	gear, err := h.service.FindByID(r.Context(), idParam)

	if err != nil {
		utils.Error(w, http.StatusNotFound, "Gear not found")

		return
	}

	utils.JSON(w, http.StatusOK, "Gear found", gear)
}

func (h *GearHandler) Update(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid UUID")

		return
	}

	var gear dto.Gear

	if err := json.NewDecoder(r.Body).Decode(&gear); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")

		return
	}

	gear.ID = id

	err = h.service.Update(r.Context(), &gear)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to update gear")

		return
	}

	utils.JSON(
		w,
		http.StatusOK,
		"Gear updated successfully",
		gear,
	)
}

func (h *GearHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.service.Delete(r.Context(), id)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Failed to delete gear")
		return
	}

	utils.JSON(
		w,
		http.StatusOK,
		"Gear deleted successfully",
		nil,
	)
}
