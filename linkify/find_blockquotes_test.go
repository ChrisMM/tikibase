package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBlockquotes(t *testing.T) {
	give := "foo bar\n\n    ```\n    text inside the blockquote\n    ```\nmore text"
	have := findBlockquotes(give)
	want := []string{"\n    ```\n    text inside the blockquote\n    ```"}
	assert.Equal(t, want, have)
}
