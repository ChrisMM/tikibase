package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
)

type containsStringTestData struct {
	desc     string
	haystack string
	needles  []string
	want     bool
}

func TestContainsString(t *testing.T) {
	tests := []containsStringTestData{
		{desc: "empty", haystack: "", needles: []string{}, want: true},
		{desc: "single match", haystack: "one two three", needles: []string{"two"}, want: true},
		{desc: "multi match", haystack: "one two three", needles: []string{"two", "one"}, want: true},
		{desc: "partial match", haystack: "one two three", needles: []string{"ree"}, want: true},
		{desc: "no match", haystack: "one two", needles: []string{"three"}, want: false},
		{desc: "different case", haystack: "One Two", needles: []string{"two"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if helpers.ContainsStrings(tt.haystack, tt.needles) != tt.want {
				t.Fatalf("expected %v to be inside %s", tt.needles, tt.haystack)
			}
		})
	}
}
