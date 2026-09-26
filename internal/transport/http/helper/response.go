package helper

import (
	"encoding/json"
	"errors"
	"go-auth-service/internal/domain"
	"math"

	"net/http"
	"time"
)

type envelope struct {
	Success   bool         `json:"isSuccess"`
	Message   string       `json:"message"`
	Data      any          `json:"data,omitempty"`
	Errors    any          `json:"errors,omitempty"`
	ErrorCode string       `json:"errorCode,omitempty"`
	Meta      responseMeta `json:"meta"`
}

type responseMeta struct {
	Pagination *paginationMeta `json:"pagination,omitempty"`
	Timestamp  string          `json:"timestamp"`
}

type paginationMeta struct {
	CurrentPage int  `json:"currentPage"`
	PerPage     int  `json:"perPage"`
	TotalItems  int  `json:"totalItems"`
	TotalPages  int  `json:"totalPages"`
	HasNextPage bool `json:"hasNextPage"`
	HasPrevPage bool `json:"hasPrevPage"`
}


func Respond(w http.ResponseWriter, status int, message string, data any, err error) {
	if err == nil {
		writeJSON(w, status, message, data)
		return
	}

	switch {

	// Handler/controller error
	case errors.Is(err, ErrBadParam):
		writeError(w, http.StatusBadRequest, err.Error(), "BAD_REQUEST")
	case errors.Is(err, ErrValidation):
		writeError(w, http.StatusBadRequest, err.Error(), "BAD_REQUEST")
		fields, ok := ValidationFields(err)
		if !ok {
			fields = []FieldError{{Field: "request", Message: "Request is invalid"}}
		}
		writeValidationError(w, http.StatusBadRequest, "Validation failed", fields)

	// Domain/service error
	case errors.Is(err, domain.ErrConflict):
		writeError(w, http.StatusConflict, err.Error(), "CONFLICT")
	case errors.Is(err, domain.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error(), "NOT_FOUND")
	case errors.Is(err, domain.ErrDuplicate):
		writeError(w, http.StatusConflict, err.Error(), "DUPLICATE_ENTRY")

	default:
		writeError(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_SERVER_ERROR")
	}
}


func writeJSON(w http.ResponseWriter, status int, message string, data any) {
	writeResponse(w, status, envelope{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    newMeta(nil),
	})
}

func writePage(w http.ResponseWriter, status int, message string, data any, page, size, total int) {
	writeResponse(w, status, envelope{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    newMeta(newPaginationMeta(page, size, total)),
	})
}

func writeValidationError(w http.ResponseWriter, status int, message string, errors any) {
	writeResponse(w, status, envelope{
		Success: false,
		Message: message,
		Errors:  errors,
		Meta:    newMeta(nil),
	})
}

func writeError(w http.ResponseWriter, status int, message, errorCode string) {
	writeResponse(w, status, envelope{
		Success:   false,
		Message:   message,
		ErrorCode: errorCode,
		Meta:      newMeta(nil),
	})
}

func writeResponse(w http.ResponseWriter, status int, response envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

func newMeta(pagination *paginationMeta) responseMeta {
	return responseMeta{
		Pagination: pagination,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}
}

func newPaginationMeta(page, size, total int) *paginationMeta {
	totalPages := int(math.Ceil(float64(total) / float64(size)))
	if total == 0 {
		totalPages = 0
	}

	return &paginationMeta{
		CurrentPage: page,
		PerPage:     size,
		TotalItems:  total,
		TotalPages:  totalPages,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}
}
