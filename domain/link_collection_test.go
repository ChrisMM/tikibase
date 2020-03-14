package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestLinkCollection_HasTarget(t *testing.T) {
	links := domain.ScaffoldLinkCollection([]string{"one.md", "two.md"})
	assert.True(t, links.HasTarget("one.md"))
	assert.True(t, links.HasTarget("two.md"))
	assert.False(t, links.HasTarget("zonk.md"))
}
