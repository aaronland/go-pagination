package cursor

import (
	_ "fmt"
	"github.com/aaronland/go-pagination"
	"net/url"
)

type CursorResults struct {
	pagination.Results `json:",omitempty"`
	TotalCount            int64   `json:"total"`
	PerPageCount          int64   `json:"per_page"`
	PageCount             int64   `json:"page"`
	PagesCount            int64   `json:"pages"`
	NextPageURI           int64   `json:"next_page"`
	PreviousPageURI       int64   `json:"previous_page"`
	PagesRange            []int64 `json:"pages_range"`
	CursorNext            string  `json:"next_cursor"`
	CursorPrevious            string  `json:"previous_cursor"`	
}

func (p *CursorResults) Method() pagination.Method {
	return pagination.Cursor
}

func (p *CursorResults) Total() int64 {
	return -1
}

func (p *CursorResults) Cursor() string {
	return p.PageCursor
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

func (p *CursorResults) NextURL(u *url.URL) string {

	cursor := p.Cursor()

	if cursor == "" {
		return "#"
	}

	q := u.Query()

	q.Set("cursor", cursor)
	u.RawQuery = q.Encode()

	return u.String()
}

func (p *CursorResults) PreviousURL(u *url.URL) string {
	return "#"
}

func (p *CursorResults) Range() []int64 {
	return p.PagesRange
}

func NewPaginationFromCursor(cursor string) (pagination.Results, error) {

	pg := new(CursorResults)
	pg.PageCursor = cursor

	return pg, nil
}
