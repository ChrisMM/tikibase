package mentions_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/test"

	"github.com/kevgo/tikibase/mentions"
)

func TestLinksToDocs(t *testing.T) {
	// create documents
	_, dc := test.NewDocumentCreator(t)
	doc1 := dc.CreateDocument("one.md", "# The one\n")
	doc2 := dc.CreateDocument("two.md", "# The other\n")
	doc3 := dc.CreateDocument("three.md", "# The third\n")

	// convert links
	links := []domain.TikiLink{
		domain.NewTikiLink("1-2", doc1.TitleSection(), doc2),
		domain.NewTikiLink("1-3", doc1.TitleSection(), doc3),
		domain.NewTikiLink("2-1", doc2.TitleSection(), doc1),
		domain.NewTikiLink("2-3", doc2.TitleSection(), doc3),
		domain.NewTikiLink("3-1", doc3.TitleSection(), doc1),
		domain.NewTikiLink("3-2", doc3.TitleSection(), doc2),
	}
	actual := mentions.LinksToDocs(links)

	// verify
	expected := map[domain.TikiDocumentFilename][]domain.TikiLink{
		domain.TikiDocumentFilename("one.md"):   {links[2], links[4]},
		domain.TikiDocumentFilename("two.md"):   {links[0], links[5]},
		domain.TikiDocumentFilename("three.md"): {links[1], links[3]},
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(links[0], doc1.TitleSection(), doc1))
	if diff != "" {
		t.Fatal(diff)
	}
}
