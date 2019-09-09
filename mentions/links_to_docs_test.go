package mentions_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kevgo/tikibase/domain"

	"github.com/kevgo/tikibase/mentions"
	"github.com/kevgo/tikibase/test"
)

func TestLinksToDocs(t *testing.T) {
	tb := test.NewTempDirectoryTikiBase(t)
	doc1, err := tb.CreateDocument("one.md", "# The one\n[the other](two.md)")
	if err != nil {
		t.Fatal(err)
	}
	doc2, err := tb.CreateDocument("two.md", "# The other\n[the one](one.md)")
	if err != nil {
		t.Fatal(err)
	}
	links, err := tb.TikiLinks()
	if err != nil {
		t.Fatal(err)
	}

	actual := mentions.LinksToDocs(links)
	expected := map[domain.TikiDocumentFilename][]domain.TikiLink{
		domain.TikiDocumentFilename("one.md"): {
			domain.NewTikiLink("the one", doc2.TitleSection(), doc1),
		},
		domain.TikiDocumentFilename("two.md"): {
			domain.NewTikiLink("the other", doc1.TitleSection(), doc2),
		},
	}
	diff := cmp.Diff(expected, actual, cmp.AllowUnexported(links[0], doc1.TitleSection(), doc1))
	if diff != "" {
		t.Fatal(diff)
	}
}
