package test

import (
	"io/ioutil"
	"testing"

	"github.com/kevgo/tikibase/domain"
)

// NewTempTikiBase provides an empty TikiBase instance
// in the system's temp directory.
// Repeated calls to this return unique instances.
func NewTempTikiBase(t *testing.T) *domain.TikiBase {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("cannot create temp dir")
	}
	result, err := domain.NewTikiBase(tmpDir)
	if err != nil {
		t.Fatalf("cannot create new TikiBase: %v", err)
	}
	return result
}
