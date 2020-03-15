package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestLinks_HasLinkTo(t *testing.T) {
	links := domain.ScaffoldLinks([]string{"one.md", "two.md"})
	assert.True(t, links.HasLinkTo("one.md"))
	assert.True(t, links.HasLinkTo("two.md"))
	assert.False(t, links.HasLinkTo("zonk.md"))
}
