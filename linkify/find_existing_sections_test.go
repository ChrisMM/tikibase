package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindExistingSections(t *testing.T) {
	text := "# One\n\nThe one.\n\n### what is it\n\n### what does it\n\n"
	have := findExistingSections(text)
	want := []string{
		"# One\n",
		"### what does it\n",
		"### what is it\n",
	}
	assert.Equal(t, want, have)
}
