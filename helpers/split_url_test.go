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
	testData := map[string]result{
		"one.md":            {"one.md", ""},
		"one.md#what-is-it": {"one.md", "what-is-it"},
	}
	for input, expected := range testData {
		filename, anchor := helpers.SplitURL(input)
		if filename != expected.filename {
			t.Fatalf("mismatching filename for '%s': expected '%s', got '%s'", input, expected.filename, filename)
		}
		if anchor != expected.anchor {
			t.Fatalf("mismatching anchor for '%s': expected '%s', got '%s'", input, expected.anchor, anchor)
		}
	}
}
