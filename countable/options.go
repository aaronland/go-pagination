package countable

import (
	"github.com/aaronland/go-pagination"
)

const PER_PAGE int64 = 10
const PAGE int64 = 1
const SPILL int64 = 2
const COLUMN string = "*"

type CountablePaginatedOptions struct {
	pagination.PaginatedOptions
	per_page int64
	page     int64
	spill    int64
	column   string
}

func DefaultCountablePaginatedOptions() (pagination.PaginatedOptions, error) {
	return NewCountablePaginatedOptions(PER_PAGE, PAGE, SPILL, COLUMN)
}

func NewCountablePaginatedOptions(per_page int64, page int64, spill int64, column string) (pagination.PaginatedOptions, error) {

	opts := CountablePaginatedOptions{
		per_page: per_page,
		page:     page,
		spill:    spill,
		column:   column,
	}

	return &opts, nil
}

func (o *CountablePaginatedOptions) PerPage() int64 {
	return o.per_page
}

func (o *CountablePaginatedOptions) Page() int64 {
	return o.page
}

func (o *CountablePaginatedOptions) Spill() int64 {
	return o.spill
}

func (o *CountablePaginatedOptions) Column() string {
	return o.column
}
