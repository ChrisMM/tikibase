package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
)

func TestIsURL(t *testing.T) {
	testData := map[string]bool{
		"one.md":             false,
		"http://google.com":  true,
		"https://google.com": true,
		"//google.com":       true,
	}
	for input, expected := range testData {
		actual := helpers.IsURL(input)
		if actual != expected {
			t.Fatalf("expected '%s' to yield %t but got %t", input, expected, actual)
		}
	}
}
