package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
	"github.com/stretchr/testify/assert"
)

func TestReverseStringList(t *testing.T) {
	strings := []string{"one", "two"}
	helpers.ReverseStringList(strings)
	assert.Equal(t, []string{"two", "one"}, strings)
}
