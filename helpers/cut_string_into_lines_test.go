package helpers_test

import (
	"reflect"
	"testing"

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
			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("expected %q, got %q", expected, actual)
			}
		})
	}
}
