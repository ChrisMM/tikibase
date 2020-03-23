package linkify_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/linkify"
	"github.com/stretchr/testify/assert"
)

func TestDocLinks(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# The One\n"},
		{FileName: "two.md", Content: "# The Other\n"},
	})
	have, err := linkify.DocLinks(docs)
	assert.Nil(t, err)
	want := map[string]string{
		"one":       "one.md",
		"The One":   "one.md",
		"The Other": "two.md",
		"two":       "two.md",
	}
	assert.Equal(t, want, have)
}
