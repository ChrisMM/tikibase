package check

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestDuplicateTargets(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "# One\n\n### alpha\n\ntext here\n\n### alpha\n\ntext\n"},
		{FileName: "2.md", Content: "# Two\n\n### gamma\n\n### gamma\n"},
	})
	duplicates, err := findDuplicateTargets(docs)
	assert.Nil(t, err)
	assert.EqualValues(t, []string{"1.md#alpha", "2.md#gamma"}, duplicates)
}
