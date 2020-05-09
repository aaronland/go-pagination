package pagination

import (
	"net/url"
)

type PaginatedOptions interface {
	PerPage(...int) int
	Page(...int) int
	Spill(...int) int
	Column(...string) string
}

type PaginatedResponse interface {
	Rows() []interface{}
	Pagination() Pagination
}

type PaginatedResponseCallback func(PaginatedResponse) error

type Pagination interface {
	Total() int
	PerPage() int
	Page() int
	Pages() int
	NextPage() int
	PreviousPage() int
	NextURL(u *url.URL) string
	PreviousURL(u *url.URL) string
	Range() []int
}
