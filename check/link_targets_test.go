package check

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindLinkTargets(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "# One\n\n### alpha\n\ntext here\n\n### beta\n\ntext\n"},
		{FileName: "2.md", Content: "# Two\n\n### gamma\n\n### delta\n"},
	})
	resources := domain.ScaffoldResourceFiles([]string{"photo.jpg"})
	ltc, err := findLinkTargets(docs, resources)
	assert.Nil(t, err)
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

func TestLinkTargets(t *testing.T) {
	ltc := make(linkTargets)
	ltc.Add("1.md")
	ltc.Add("1.md#foo")
	assert.True(t, ltc.Contains("1.md"))
	assert.True(t, ltc.Contains("1.md#foo"))
	assert.False(t, ltc.Contains("1.md#bar"))
	assert.False(t, ltc.Contains("2.md"))
}

func TestLinkTargets_AddDuplicate(t *testing.T) {
	ltc := make(linkTargets)
	ltc.Add("1.md")
	ltc.Add("1.md")
}

func TestLinkTargets_String(t *testing.T) {
	ltc := linkTargets{}
	ltc.Add("one")
	ltc.Add("two")
	assert.Equal(t, "[one, two]", ltc.String())
}

func assertContains(t *testing.T, targets linkTargets, target string) {
	if _, exists := targets[target]; !exists {
		t.Errorf("expected %q in %s", target, targets)
	}
}
