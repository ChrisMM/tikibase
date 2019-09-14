package helpers_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/helpers"
)

func TestCutStringIntoLines(t *testing.T) {
	expected := []string{"one\n", "two\n", "three\n"}
	actual := helpers.CutStringIntoLines("one\ntwo\nthree\n")
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatalf("test 1: expected '%s', got '%s'", expected, actual)
	}

	expected = []string{"one\n", "two\n", "three"}
	actual = helpers.CutStringIntoLines("one\ntwo\nthree")
	diff = cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatalf("test 2: expected '%s', got '%s'", expected, actual)
	}

	expected = []string{""}
	actual = helpers.CutStringIntoLines("")
	diff = cmp.Diff(expected, actual)
	if diff != "" {
		t.Fatalf("test 3: expected '%s', got '%s'", expected, actual)
	}
}
