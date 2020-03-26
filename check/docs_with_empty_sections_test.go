package check

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestDocsWithEmptySections_EmptyTitle(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "#  \nfoo"},
	})
	have, err := docsWithEmptySections(docs)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.md"}, have)
}

func TestDocsWithEmptySections_EmptySection(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "# One\n###  \nfoo"},
	})
	have, err := docsWithEmptySections(docs)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1.md"}, have)
}
