package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
)

type result struct {
	filename string
	anchor   string
}

func TestSplitURL(t *testing.T) {
	tests := map[string]result{
		"one.md":            {"one.md", ""},
		"one.md#what-is-it": {"one.md", "what-is-it"},
	}
	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			filename, anchor := helpers.SplitURL(input)

			if filename != expected.filename {
				t.Fatalf("mismatching filename: expected %q, got %q", expected.filename, filename)
			}
			if anchor != expected.anchor {
				t.Fatalf("mismatching anchor: expected %q, got %q", expected.anchor, anchor)
			}
		})
	}
}
