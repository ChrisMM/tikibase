package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
)

func TestIsURL(t *testing.T) {
	tests := map[string]bool{
		"one.md":             false,
		"http://google.com":  true,
		"https://google.com": true,
		"//google.com":       true,
	}
	for input, expected := range tests {
		t.Run(input, func(t *testing.T) {
			actual := helpers.IsURL(input)
			if actual != expected {
				t.Fatalf("expected %q to yield %t but got %t", input, expected, actual)
			}
		})
	}
}
