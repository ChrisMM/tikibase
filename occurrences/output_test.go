package occurrences_test

import (
	"regexp"
	"testing"

	"github.com/kevgo/tikibase/occurrences"
)

func TestOutput_Footer_complete(t *testing.T) {
	output := occurrences.NewDotOutput()
	output.Created()
	output.Updated()
	output.Updated()
	output.Deleted()
	output.Deleted()
	output.Deleted()
	expected := regexp.MustCompile(`1 created, 2 updated, 3 deleted in \dm?s`)
	actual := output.Footer()
	if !expected.MatchString(actual) {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestOutput_Footer_created(t *testing.T) {
	output := occurrences.NewDotOutput()
	output.Created()
	expected := regexp.MustCompile(`1 created in \dm?s`)
	actual := output.Footer()
	if !expected.MatchString(actual) {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestOutput_Footer_noChanges(t *testing.T) {
	output := occurrences.NewDotOutput()
	expected := regexp.MustCompile(`no changes, \dm?s`)
	actual := output.Footer()
	if !expected.MatchString(actual) {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}
