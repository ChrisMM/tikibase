package helpers_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/helpers"
)

func TestCutStringIntoLines(t *testing.T) {
	tests := map[string][]string{
		"one\ntwo\nthree\n": {"one\n", "two\n", "three\n"},
		"one\ntwo\nthree":   {"one\n", "two\n", "three"},
		"":                  {""},
	}
	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			actual := helpers.CutStringIntoLines(input)
			diff := cmp.Diff(expected, actual)
			if diff != "" {
				t.Fatalf("expected %q, got %q", expected, actual)
			}
		})
	}
}
