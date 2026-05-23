package errors

import (
	"strconv"
)

// Pagination represents TMF-compliant pagination parameters
type Pagination struct {
	Offset int
	Limit  int
}

func GetPagination(offsetStr, limitStr string) Pagination {
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 20 // Default limit
	}
	return Pagination{Offset: offset, Limit: limit}
}
