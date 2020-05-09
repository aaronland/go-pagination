package pagination

import (
	"net/url"
)

type PaginationOptions struct {
	PerPage   int64
	Page      int64
	Spill     int64
	Countable string
}

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
