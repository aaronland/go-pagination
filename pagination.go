package pagination

import (
	"math"
	"net/url"
)

// PaginationOptions provides an interface for criteria used to paginate a query.
type PaginationOptions interface {
	PerPage(...int64) int64
	Page(...int64) int64
	Spill(...int64) int64
	Column(...string) string
	Cursor(...string) string
}

// Pagination provides an interface for pagination information for a query response.
type Pagination interface {
	// The total number of results for a query.
	Total() int64
	// The number of results per page for a query.
	PerPage() int64
	// The current page number (offset) for a paginated query response.
	Page() int64
	// The total number of pages for a paginated query response.
	Pages() int64
	// The cursor (token) to use to advance to the next set of results in a query response.
	Cursor() string
	// The next page (offset) for a paginated query response.
	NextPage() int64
	// The previous page (offset) for a paginated query response.	
	PreviousPage() int64
	NextURL(u *url.URL) string
	PreviousURL(u *url.URL) string
	Range() []int64
}

func PagesForCount(opts PaginationOptions, total_count int64) int64 {

	per_page := int64(math.Max(1.0, float64(opts.PerPage())))
	spill := int64(math.Max(1.0, float64(opts.Spill())))

	if spill >= per_page {
		spill = per_page - 1
	}

	pages := int64(math.Ceil(float64(total_count) / float64(per_page)))
	return pages
}
