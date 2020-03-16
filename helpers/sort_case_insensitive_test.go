package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSortCaseInsensitive(t *testing.T) {
	list := []string{"one", "two", "One", "Two", "ONE"}
	helpers.SortCaseInsensitive(list)
	assert.Equal(t, list, []string{"one", "One", "ONE", "two", "Two"})
}
