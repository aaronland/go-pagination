package cursor

import (
	"github.com/aaronland/go-pagination"
	"testing"
)

func TestCursorOptions(t *testing.T) {

	opts, err := NewCursorOptions()

	if err != nil {
		t.Fatalf("Failed to create new pagination options, %v", err)
	}

	if opts.Method() != pagination.Cursor {
		t.Fatalf("Invalid method")
	}
}
