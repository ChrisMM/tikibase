package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSections(t *testing.T) {
	text := "# One\n\nThe one.\n\n### what is it\n\n### what does it\n\n"
	have := findSections(text)
	want := []string{
		"# One\n",
		"### what does it\n",
		"### what is it\n",
	}
	assert.Equal(t, want, have)
}
