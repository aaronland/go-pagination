package pagination

import (
	"github.com/jtacoma/uritemplates"
)

type Method uint8

const (
	Countable Method = iota
	Cursor
)

// Options provides a common interface for criteria used to paginate a query.
type Options interface {
	// The number of results to return in a Results
	PerPage(...int64) int64
	// The current page to return query results for
	Page(...int64) int64
	// An optional number of extra query results to return for
	Spill(...int64) int64
	// The name of the column to use when paginating results
	Column(...string) string
	NextCursor(...string) string
	PreviousCursor(...string) string
	Method() Method
}

// Results provides an interface for pagination information for a query response.
type Results interface {
	// The total number of results for a query.
	Total() int64
	// The number of results per page for a query.
	PerPage() int64
	// The current page number (offset) for a paginated query response.
	Page() int64
	// The total number of pages for a paginated query response.
	Pages() int64
	// The cursor (token) to use to advance to the next set of results in a query response.
	NextCursor() string
	// The cursor (token) to use to rewind to the previous set of results in a query response.
	PreviousCursor() string
	// The next page (offset) for a paginated query response.
	NextPage() int64
	// The previous page (offset) for a paginated query response.
	PreviousPage() int64
	// The URL to the next set of results in a query response
	NextURL(t *uritemplates.UriTemplate) (string, error)
	// The URL to the previous set of results in a query response
	PreviousURL(*uritemplates.UriTemplate) (string, error)
	Method() Method
}
