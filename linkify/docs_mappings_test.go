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
	assert.Equal(t, "[The One](one.md)", have[0].replaceWith)
	assert.Equal(t, "one.md", have[0].file)
}
