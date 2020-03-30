package linkify

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestDocsMappings(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "one.md", Content: "# The One\n"},
	})
	have, err := docsMappings(docs)
	assert.Nil(t, err)
	want := []docMapping{
		{name: "The One", file: "one.md"},
		{name: "one", file: "one.md"},
	}
	assert.Equal(t, want, have)
}
