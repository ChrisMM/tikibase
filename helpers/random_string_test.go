package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	have := helpers.RandomString(10)
	assert.Len(t, have, 10)
}
