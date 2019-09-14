package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

func TestSectionCollectionReplace(t *testing.T) {

}

func TestSectionCollectionText(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	actual := sections.Text()
	expected := "section 1\nsection 2\n"
	if actual != expected {
		t.Fatalf("mismatching content: expected '%s', got '%s'", expected, actual)
	}
}
