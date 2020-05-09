package pagination

import (
	"net/url"
)

type PaginatedOptions interface {
	PerPage() int64
	Page() int64
	Spill() int64
	Column() string
}

type PaginatedResponse interface {
	Rows() []interface{}
	Pagination() Pagination
}

type PaginatedResponseCallback func(PaginatedResponse) error

type Pagination interface {
	Total() int64
	PerPage() int64
	Page() int64
	Pages() int64
	NextPage() int64
	PreviousPage() int64
	NextURL(u *url.URL) string
	PreviousURL(u *url.URL) string
	Range() []int64
}
