package cursor

import (
	"fmt"
	"github.com/aaronland/go-pagination"
	"github.com/jtacoma/uritemplates"
)

// type CursorResults implements the pagination.Results interface for cursor or token-based pagination.
type CursorResults struct {
	pagination.Results `json:",omitempty"`
	TotalCount         int64   `json:"total"`
	PerPageCount       int64   `json:"per_page"`
	PageCount          int64   `json:"page"`
	PagesCount         int64   `json:"pages"`
	NextPageURI        int64   `json:"next_page"`
	PreviousPageURI    int64   `json:"previous_page"`
	PagesRange         []int64 `json:"pages_range"`
	CursorNext         string  `json:"next_cursor"`
	CursorPrevious     string  `json:"previous_cursor"`
}

func (p *CursorResults) Method() pagination.Method {
	return pagination.Cursor
}

func (p *CursorResults) Total() int64 {
	return -1
}

func (p *CursorResults) NextCursor() string {
	return p.CursorNext
}

func (p *CursorResults) PreviousCursor() string {
	return p.CursorPrevious
}

func (p *CursorResults) PerPage() int64 {
	return p.PerPageCount
}

func (p *CursorResults) Page() int64 {
	return -1
}

func (p *CursorResults) Pages() int64 {
	return -1
}

func (p *CursorResults) NextPage() int64 {
	return -1
}

func (p *CursorResults) PreviousPage() int64 {
	return -1
}

func (p *CursorResults) NextURL(t *uritemplates.UriTemplate) (string, error) {

	cursor := p.NextCursor()

	if cursor == "" {
		return "#", nil
	}

	values := map[string]interface{}{
		"next": cursor,
	}

	uri, err := t.Expand(values)

	if err != nil {
		return "", fmt.Errorf("Failed to expand URI template, %w", err)
	}

	return uri, nil
}

func (p *CursorResults) PreviousURL(t *uritemplates.UriTemplate) (string, error) {

	cursor := p.PreviousCursor()

	if cursor == "" {
		return "#", nil
	}

	values := map[string]interface{}{
		"previous": cursor,
	}

	uri, err := t.Expand(values)

	if err != nil {
		return "", fmt.Errorf("Failed to expand URI template, %w", err)
	}

	return uri, nil
}

func NewPaginationFromCursors(previous string, next string) (pagination.Results, error) {

	pg := new(CursorResults)
	pg.CursorPrevious = previous
	pg.CursorNext = next

	return pg, nil
}
