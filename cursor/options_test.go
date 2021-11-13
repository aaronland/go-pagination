package cursor

import (
	"testing"
)

func TestCursorPaginationOptions(t *testing.T) {

	opts, err := NewCursorPaginationOptions()

	if err != nil {
		t.Fatalf("Failed to create new pagination options, %v", err)
	}
}
