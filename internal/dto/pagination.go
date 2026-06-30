package dto

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

type PaginatedGearsResponse struct {
	Data []Gear         `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
