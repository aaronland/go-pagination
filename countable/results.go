package countable

import (
	"fmt"
	"github.com/aaronland/go-pagination"
	"github.com/jtacoma/uritemplates"
	"math"
)

type CountableResults struct {
	pagination.Results `json:",omitempty"`
	TotalCount         int64 `json:"total"`
	PerPageCount       int64 `json:"per_page"`
	PageCount          int64 `json:"page"`
	PagesCount         int64 `json:"pages"`
	NextPageURI        int64 `json:"next_page"`
	PreviousPageURI    int64 `json:"previous_page"`
}

func (p *CountableResults) Method() pagination.Method {
	return pagination.Countable
}

func (p *CountableResults) Total() int64 {
	return p.TotalCount
}

func (p *CountableResults) PerPage() int64 {
	return p.PerPageCount
}

func (p *CountableResults) Page() int64 {
	return p.PageCount
}

func (p *CountableResults) Pages() int64 {
	return p.PagesCount
}

func (p *CountableResults) NextPage() int64 {
	return p.NextPageURI
}

func (p *CountableResults) PreviousPage() int64 {
	return p.PreviousPageURI
}

func (p *CountableResults) NextCursor() string {
	return ""
}

func (p *CountableResults) PreviousCursor() string {
	return ""
}

func (p *CountableResults) NextURL(t *uritemplates.UriTemplate) (string, error) {

	next := p.NextPage()

	if next == 0 {
		return "#", nil
	}
	values := map[string]interface{}{
		"next": next,
	}

	uri, err := t.Expand(values)

	if err != nil {
		return "", fmt.Errorf("Failed to expand URI template, %w", err)
	}

	return uri, nil
}

func (p *CountableResults) PreviousURL(t *uritemplates.UriTemplate) (string, error) {

	previous := p.PreviousPage()

	if previous == 0 {
		return "#", nil
	}

	values := map[string]interface{}{
		"previous": previous,
	}

	uri, err := t.Expand(values)

	if err != nil {
		return "", fmt.Errorf("Failed to expand URI template, %w", err)
	}

	return uri, nil
}

func NewResultsFromCount(total_count int64) (pagination.Results, error) {

	opts, err := NewCountableOptions()

	if err != nil {
		return nil, err
	}

	return NewResultsFromCountWithOptions(opts, total_count)
}

func NewResultsFromCountWithOptions(opts pagination.Options, total_count int64) (pagination.Results, error) {

	page := int64(math.Max(1.0, float64(opts.Page())))
	per_page := int64(math.Max(1.0, float64(opts.PerPage())))

	pages := PagesForCount(opts, total_count)

	next_page := int64(0)
	previous_page := int64(0)

	if pages > 1 {

		if page > 1 {
			previous_page = page - 1

		}

		if page < pages {
			next_page = page + 1
		}

	}

	pages_range := make([]int64, 0)

	var range_min int64
	var range_max int64
	var range_mid int64

	var rfloor int64
	var adjmin int64
	var adjmax int64

	if pages > 10 {

		range_mid = 7
		rfloor = int64(math.Floor(float64(range_mid) / 2.0))

		range_min = page - rfloor
		range_max = page + rfloor

		if range_min <= 0 {

			adjmin = int64(math.Abs(float64(range_min)))

			range_min = 1
			range_max = page + adjmin + 1
		}

		if range_max >= pages {

			adjmax = range_max - pages

			range_min = range_min - adjmax
			range_max = pages
		}

		for i := range_min; range_min <= range_max; range_min++ {
			pages_range = append(pages_range, i)
		}
	}

	pg := &CountableResults{
		TotalCount:      total_count,
		PerPageCount:    per_page,
		PageCount:       page,
		PagesCount:      pages,
		NextPageURI:     next_page,
		PreviousPageURI: previous_page,
	}

	return pg, nil
}
