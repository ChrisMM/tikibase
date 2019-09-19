package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

func TestScaffoldTikiLink(t *testing.T) {
	actual := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{})
	if actual.Title() == "" {
		t.Fatal("link scaffolded wrong")
	}
}

func TestTikiLinkSourceSection(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{})
	link := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{SourceSection: &section})
	if link.SourceSection() != &section {
		t.Fatalf("wrong section returned")
	}
}

func TestTikiLinkTargetDocument(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	link := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{TargetDocument: &doc})
	if link.TargetDocument() != &doc {
		t.Fatalf("wrong TargetDocument returned")
	}
}

func TestTikiLinkTitle(t *testing.T) {
	expected := "My Title"
	link := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{Title: expected})
	if link.Title() != expected {
		t.Fatalf("wrong TargetDocument returned")
	}
}
