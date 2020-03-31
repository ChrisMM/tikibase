package linkify

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceholders(t *testing.T) {
	placeholders := NewPlaceholders()
	p1 := placeholders.Make("foo")
	p2 := placeholders.Make("bar")
	assert.NotEqual(t, p1, p2)
	phrases := []string{}
	for phrase := range placeholders {
		phrases = append(phrases, phrase)
	}
	sort.Strings(phrases)
	assert.Equal(t, []string{"bar", "foo"}, phrases)
}
