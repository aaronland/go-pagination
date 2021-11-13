package cursor

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCursorPagination(t *testing.T) {

	cursor := "next"

	pg, err := NewPaginationFromCursor(cursor)

	if err != nil {
		t.Fatalf("Failed to create cursor, %v", err)
	}

	if pg.Cursor() != cursor {
		t.Fatalf("Invalid cursor")
	}

	if pg.Pages() != -1 {
		t.Fatalf("Invalid pages")
	}

	u, err := url.Parse("http://example.com")

	if err != nil {
		t.Fatalf("Failed to parse URL, %v", err)
	}

	next_url := pg.NextURL(u)

	expected_url := fmt.Sprintf("http://example.com?cursor=%s", cursor)

	if next_url != expected_url {
		t.Fatalf("Unexpected URL")
	}
}
