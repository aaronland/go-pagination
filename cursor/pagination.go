package cursor

import (
	"fmt"
	"github.com/aaronland/go-pagination"
	"net/url"
)

type CursorPagination struct {
	pagination.Pagination `json:",omitempty"`
	TotalCount            int64   `json:"total"`
	PerPageCount          int64   `json:"per_page"`
	PageCount             int64   `json:"page"`
	PagesCount            int64   `json:"pages"`
	NextPageURI           int64   `json:"next_page"`
	PreviousPageURI       int64   `json:"previous_page"`
	PagesRange            []int64 `json:"pages_range"`
	PageCursor            string  `json:"cursor"`
}

func (p *CursorPagination) Total() int64 {
	return -1
}

func (p *CursorPagination) Cursor() string {
	return p.PageCursor
}

func (p *CursorPagination) PerPage() int64 {
	return p.PerPageCount
}

func (p *CursorPagination) Page() int64 {
	return -1
}

func (p *CursorPagination) Pages() int64 {
	return -1
}

func (p *CursorPagination) NextPage() int64 {
	return -1
}

func (p *CursorPagination) PreviousPage() int64 {
	return -1
}

func (p *CursorPagination) NextURL(u *url.URL) string {

	next := p.NextPage()

	if next == 0 {
		return "#"
	}

	q := u.Query()

	q.Set("page", fmt.Sprintf("%d", next))
	u.RawQuery = q.Encode()

	return u.String()
}

func (p *CursorPagination) PreviousURL(u *url.URL) string {

	previous := p.PreviousPage()

	if previous == 0 {
		return "#"
	}

	q := u.Query()

	q.Set("page", fmt.Sprintf("%d", previous))
	u.RawQuery = q.Encode()

	return u.String()
}

func (p *CursorPagination) Range() []int64 {
	return p.PagesRange
}

func NewPaginationFromCursor(cursor string) (pagination.Pagination, error) {

	pg := new(CursorPagination)
	pg.PageCursor = cursor

	return pg, nil
}
