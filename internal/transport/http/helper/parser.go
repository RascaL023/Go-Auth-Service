package helper

import (
	"net/http"
	"strconv"
)

func ParseID(rawID string) (int64, error) {
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return 0, NewError(
			ErrBadParam,
			"Invalid ID",
		)
	}

	if id <= 0 {
		return 0, NewError(
			ErrBadParam,
			"ID must be greater than zero",
		)
	}

	return id, nil
}

func Pagination(r *http.Request) (int, int) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if page <= 0 { page = 1 }
	if size <= 0 { size = 10 }
	if size > 100 { size = 100 }

	return page, size
}
