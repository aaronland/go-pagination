// package countable provides implementions of the pagintion.Options and pagination.Results interfaces for use with page-based or numbered pagination.
package countable

import (
	"github.com/aaronland/go-pagination"
	"math"
)

func NextPage(r pagination.Results) int64 {

	if r.Method() != pagination.Countable {
		return 0
	}

	return r.Next().(int64)
}

func PreviousPage(r pagination.Results) int64 {

	if r.Method() != pagination.Countable {
		return 0
	}

	return r.Previous().(int64)
}

// PagesForCount returns the number of pages that total_count will span using criteria defined in opts.
func PagesForCount(opts pagination.Options, total_count int64) int64 {

	per_page := int64(math.Max(1.0, float64(opts.PerPage())))
	spill := int64(math.Max(1.0, float64(opts.Spill())))

	if spill >= per_page {
		spill = per_page - 1
	}

	pages := int64(math.Ceil(float64(total_count) / float64(per_page)))
	return pages
}
