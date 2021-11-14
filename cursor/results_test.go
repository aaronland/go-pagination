package cursor

import (
	"fmt"
	"github.com/jtacoma/uritemplates"
	"testing"
)

func TestCursorPagination(t *testing.T) {

	previous_cursor := ""
	next_cursor := "next"

	pg, err := NewPaginationFromCursors(previous_cursor, next_cursor)

	if err != nil {
		t.Fatalf("Failed to create cursor, %v", err)
	}

	if NextCursor(pg) != next_cursor {
		t.Fatalf("Invalid cursor")
	}

	if pg.Pages() != -1 {
		t.Fatalf("Invalid pages")
	}

	next_t, err := uritemplates.Parse("http://example.com?cursor={next}")

	if err != nil {
		t.Fatalf("Failed to compile URI template, %v", err)
	}

	next_url, err := pg.NextURL(next_t)

	if err != nil {
		t.Fatalf("Failed to derive next URL, %v", err)
	}

	expected_url := fmt.Sprintf("http://example.com?cursor=%s", next_cursor)

	if next_url != expected_url {
		t.Fatalf("Unexpected URL")
	}
}
