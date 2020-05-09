package countable

import (
	"github.com/aaronland/go-pagination"
)

type CountablePaginatedResponse struct {
	pagination.PaginatedResponse
	rows       []interface{}
	pagination pagination.Pagination
}

func (r *CountablePaginatedResponse) Rows() []interface{} {
	return r.rows
}

func (r *CountablePaginatedResponse) Pagination() pagination.Pagination {
	return r.pagination
}
