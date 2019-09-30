package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestScaffoldTikiLink(t *testing.T) {
	actual := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{})
	assert.NotEqual(t, "", actual.Title())
}

func TestTikiLinkSourceSection(t *testing.T) {
	section := domain.ScaffoldSection(domain.SectionScaffold{})
	link := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{SourceSection: section})
	assert.Same(t, section, link.SourceSection())
}

func TestTikiLinkTargetDocument(t *testing.T) {
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{})
	link := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{TargetDocument: doc})
	assert.Same(t, doc, link.TargetDocument())
}

func TestTikiLinkTitle(t *testing.T) {
	text := "My Title"
	link := domain.ScaffoldTikiLink(domain.TikiLinkScaffold{Title: text})
	assert.Equal(t, text, link.Title())
}
