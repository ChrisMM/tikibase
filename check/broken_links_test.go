package check

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestIsBrokenLink(t *testing.T) {
	targets := scaffoldLinkTargets([]string{"1.md"})
	assert.False(t, isBrokenLink("1.md", "foo.md", targets))
	assert.True(t, isBrokenLink("2.md", "foo.md", targets))
}

func TestFindBrokenLinks(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "1.md", Content: "# One\n[two](2.md)\n"},
	})
	resorces := domain.ScaffoldResourceFiles([]string{"chart.png"})
	links, _ := docs.Links()
	have := findBrokenLinks(docs, resorces, links)
	want := []brokenLink{
		{Filename: "1.md", Link: "2.md"},
	}
	assert.Equal(t, want, have)
}
