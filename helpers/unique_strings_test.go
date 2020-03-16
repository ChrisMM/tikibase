package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
	"github.com/stretchr/testify/assert"
)

func TestUniqueStrings(t *testing.T) {
	assert.Equal(t, []string{"One", "one"}, helpers.UniqueStrings([]string{"one", "One", "one", "One"}))
}
