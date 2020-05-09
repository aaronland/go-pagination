package countable

import (
	"github.com/aaronland/go-pagination"
)

const PER_PAGE int64 = 10
const PAGE int64 = 1
const SPILL int64 = 2
const COUNTABLE string = "*"

func NewCountablePaginationOptions() (*pagination.PaginationOptions, error) {

	opts := &pagination.PaginationOptions{
		PerPage:   PER_PAGE,
		Page:      PAGE,
		Spill:     SPILL,
		Countable: COUNTABLE,
	}

	return opts, nil
}
