package occurrences_test

import (
	"testing"
	"time"

	"github.com/kevgo/tikibase/occurrences"
)

func TestOutput_Footer(t *testing.T) {
	tests := map[string][4]int{
		"1 created, 2 updated, 3 deleted in 45ms": {1, 2, 3, 45},
		"1 created in 45ms":                       {1, 0, 0, 45},
		"no changes, 1ms":                         {0, 0, 0, 1},
	}
	for expected, input := range tests {
		t.Run(expected, func(t *testing.T) {
			output := occurrences.ScaffoldOutput(input[0], input[1], input[2])

			actual := output.Footer(time.Duration(input[3]) * time.Millisecond)

			if actual != expected {
				t.Fatalf("expected %q, got %q", expected, actual)
			}
		})
	}
}
