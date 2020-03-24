package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
	"github.com/stretchr/testify/assert"
)

func TestDedupeStrings(t *testing.T) {
	assert.Equal(t, []string{"One", "one"}, helpers.DedupeStrings([]string{"one", "One", "one", "One"}))
}
