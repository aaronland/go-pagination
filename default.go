package pagination

import (
	"fmt"
	"net/url"
)

type DefaultPaginatedResponse struct {
	rows       []interface{}
	pagination Pagination
}

func (r *DefaultPaginatedResponse) Rows() []interface{} {
	return r.rows
}

func (r *DefaultPaginatedResponse) Pagination() Pagination {
	return r.pagination
}

type DefaultPagination struct {
	Pagination
	total         int
	per_page      int
	page          int
	pages         int
	next_page     int
	previous_page int
	pages_range   []int
}

func (p *DefaultPagination) Total() int {
	return p.total
}

func (p *DefaultPagination) PerPage() int {
	return p.per_page
}

func (p *DefaultPagination) Page() int {
	return p.page
}

func (p *DefaultPagination) Pages() int {
	return p.pages
}

func (p *DefaultPagination) NextPage() int {
	return p.next_page
}

func (p *DefaultPagination) NextURL(u *url.URL) string {

	next := p.NextPage()

	if next == 0 {
		return "#"
	}

	q := u.Query()

	q.Set("page", fmt.Sprintf("%d", next))
	u.RawQuery = q.Encode()

	return u.String()
}

func (p *DefaultPagination) PreviousURL(u *url.URL) string {

	previous := p.PreviousPage()

	if previous == 0 {
		return "#"
	}

	q := u.Query()

	q.Set("page", fmt.Sprintf("%d", previous))
	u.RawQuery = q.Encode()

	return u.String()
}

func (p *DefaultPagination) PreviousPage() int {
	return p.previous_page
}

func (p *DefaultPagination) Range() []int {
	return p.pages_range
}

type DefaultPaginatedOptions struct {
	PaginatedOptions
	per_page int
	page     int
	spill    int
	column   string
}

func (o *DefaultPaginatedOptions) PerPage(args ...int) int {

	if len(args) == 1 {
		o.per_page = args[0]
	}
	return o.per_page
}

func (o *DefaultPaginatedOptions) Page(args ...int) int {

	if len(args) == 1 {
		o.page = args[0]
	}

	return o.page
}

func (o *DefaultPaginatedOptions) Spill(args ...int) int {

	if len(args) == 1 {
		o.spill = args[0]
	}

	return o.spill
}

func (o *DefaultPaginatedOptions) Column(args ...string) string {

	if len(args) == 1 {
		o.column = args[0]
	}

	return o.column
}

func NewDefaultPaginatedOptions() PaginatedOptions {

	opts := DefaultPaginatedOptions{
		per_page: 10,
		page:     1,
		spill:    2,
		column:   "*",
	}

	return &opts
}
