package linkify_test

import (
	"testing"

	"github.com/kevgo/tikibase/linkify"
)

func TestTextContainsText(t *testing.T) {
	tests := map[string]bool{
		"one two":    true,
		"One Two":    true,
		"one\ntwo":   true,
		"one\n  two": true,
		"two one":    false,
	}
	for tt := range tests {
		if linkify.TextContainsTitle(tt, "one two") != tests[tt] {
			t.Errorf("mismatch for %q", tt)
		}
	}
}
