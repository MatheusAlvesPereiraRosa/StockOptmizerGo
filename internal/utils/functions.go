package utils

import (
	"encoding/json"
	"math"
	"net/http"

	"gear-priority-api/internal/dto"
)

func RoundMoney(value float64) float64 {
	return math.Round(value*100) / 100
}

func JSON(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(dto.Response{
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(dto.Response{
		Error: message,
	})
}

func Pagination(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func NormalizePagination(
	page int,
	limit int,
	defaultPage int,
	defaultLimit int,
	maxLimit int,
) (int, int) {
	if page < 1 {
		page = defaultPage
	}

	if limit < 1 {
		limit = defaultLimit
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}
