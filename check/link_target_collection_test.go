package check

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindLinkTargets(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "# One\n\n### alpha\n\ntext here\n\n### beta\n\ntext\n"},
		{FileName: "2.md", Content: "# Two\n\n### gamma\n\n### delta\n"},
	})
	resources := domain.ScaffoldResourceFiles([]string{"photo.jpg"})
	ltc, duplicates, err := findLinkTargets(docs, resources)
	assert.Nil(t, err)
	assert.Len(t, duplicates, 0)
	assert.Len(t, ltc, 9)
	assertContains(t, ltc, "1.md")
	assertContains(t, ltc, "1.md#one")
	assertContains(t, ltc, "1.md#alpha")
	assertContains(t, ltc, "1.md#beta")
	assertContains(t, ltc, "2.md")
	assertContains(t, ltc, "2.md#two")
	assertContains(t, ltc, "2.md#gamma")
	assertContains(t, ltc, "2.md#delta")
	assertContains(t, ltc, "photo.jpg")
}

func TestLinkTargets_duplicates(t *testing.T) {
	docs := domain.ScaffoldDocumentCollection([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "# One\n\n### alpha\n\ntext here\n\n### alpha\n\ntext\n"},
		{FileName: "2.md", Content: "# Two\n\n### gamma\n\n### gamma\n"},
	})
	files := domain.ScaffoldResourceFiles([]string{"photo.jpg"})
	_, duplicates, err := findLinkTargets(docs, files)
	assert.Nil(t, err)
	assert.EqualValues(t, []string{"1.md#alpha", "2.md#gamma"}, duplicates)
}

func TestLinkTargetCollection(t *testing.T) {
	ltc := make(linkTargetCollection)
	err := ltc.Add("1.md")
	assert.Nil(t, err)
	err = ltc.Add("1.md#foo")
	assert.Nil(t, err)
	assert.True(t, ltc.Contains("1.md"))
	assert.True(t, ltc.Contains("1.md#foo"))
	assert.False(t, ltc.Contains("1.md#bar"))
	assert.False(t, ltc.Contains("2.md"))
}

func TestLinkTargetCollection_AddDuplicate(t *testing.T) {
	ltc := make(linkTargetCollection)
	err := ltc.Add("1.md")
	assert.Nil(t, err)
	err = ltc.Add("1.md")
	assert.NotNil(t, err)
}

func TestLinkTargetCollection_String(t *testing.T) {
	ltc := linkTargetCollection{}
	err := ltc.Add("one")
	assert.Nil(t, err)
	err = ltc.Add("two")
	assert.Nil(t, err)
	assert.Equal(t, "[one, two]", ltc.String())
}

func assertContains(t *testing.T, targets linkTargetCollection, target string) {
	if _, exists := targets[target]; !exists {
		t.Errorf("expected %q in %s", target, targets)
	}
}
