package remove

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLinksToFile(t *testing.T) {
	give := "# Test\n\nHere are links to [foo](foo.md) and [bar](bar.md)!"
	have := removeLinksToFile("foo.md", give)
	want := "# Test\n\nHere are links to foo and [bar](bar.md)!"
	assert.Equal(t, want, have)
}
