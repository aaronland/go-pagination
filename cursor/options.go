package cursor

import (
	"github.com/aaronland/go-pagination"
)

const PER_PAGE int64 = 10
const PAGE int64 = 1
const SPILL int64 = 2

type CursorPaginationOptions struct {
	pagination.PaginationOptions
	perpage int64
	cursor  string
	column  string
}

func NewCursorPaginationOptions() (pagination.PaginationOptions, error) {

	opts := &CursorPaginationOptions{
		perpage: PER_PAGE,
	}

	return opts, nil
}

func (opts *CursorPaginationOptions) PerPage(args ...int64) int64 {

	if len(args) >= 1 {
		opts.perpage = args[0]
	}

	return opts.perpage
}

func (opts *CursorPaginationOptions) Cursor(args ...string) string {

	if len(args) >= 1 {
		opts.cursor = args[0]
	}

	return opts.cursor
}

func (opts *CursorPaginationOptions) Page(args ...int64) int64 {

	return 0
}

func (opts *CursorPaginationOptions) Spill(args ...int64) int64 {

	return 0
}

func (opts *CursorPaginationOptions) Column(args ...string) string {

	if len(args) >= 1 {
		opts.column = args[0]
	}

	return opts.column
}
