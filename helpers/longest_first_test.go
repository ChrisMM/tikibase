package helpers_test

import (
	"testing"

	"github.com/kevgo/tikibase/helpers"
	"github.com/stretchr/testify/assert"
)

func TestLongestFirst(t *testing.T) {
	list := []string{"wonder", "wonderland", "wonderful", "wander"}
	helpers.LongestFirst(list)
	want := []string{"wander", "wonder", "wonderful", "wonderland"}
	assert.Equal(t, want, list)
}
