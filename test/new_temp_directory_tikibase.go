package test

import (
	"io/ioutil"
	"testing"

	"github.com/kevgo/tikibase/domain"
)

// NewTempDirectoryTikiBase provides an empty DirectoryTikiBase instance
// in the system's temp directory.
// Repeated calls to this return unique instances.
func NewTempDirectoryTikiBase(t *testing.T) domain.DirectoryTikiBase {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("cannot create temp dir")
	}
	result, err := domain.NewDirectoryTikiBase(tmpDir)
	if err != nil {
		t.Fatalf("cannot create new TikiBase: %v", err)
	}
	return result
}
