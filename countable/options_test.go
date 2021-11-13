package countable

import (
	"testing"
)

func TestCountablePaginationOptions(t *testing.T) {

	_, err := NewCountablePaginationOptions()

	if err != nil {
		t.Fatalf("Failed to create new pagination options, %v", err)
	}
}
