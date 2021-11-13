package countable

import (
	"github.com/aaronland/go-pagination"
	"testing"
)

func TestCountableOptions(t *testing.T) {

	opts, err := NewCountableOptions()

	if err != nil {
		t.Fatalf("Failed to create new pagination options, %v", err)
	}

	if opts.Method() != pagination.Countable {
		t.Fatalf("Invalid method")
	}

}
