package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kevgo/tikibase/domain"
)

func TestSectionCollectionFindByTitle(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "# Title\ntext 1\n"},
		{Content: "### what is it\ntext 2\n"},
		{Content: "### what does it\ntext 3\n"},
	})
	actual, err := sections.FindByTitle("what is it")
	assert.Nil(t, err)
	assert.Same(t, sections[1], actual, "found wrong document")
}

func TestSectionCollectionRemove(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	assert.Equal(t, domain.SectionCollection{sections[0]}, sections.Remove(sections[1]))
}

func TestSectionCollectionReplace(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	newSection2 := domain.ScaffoldSection(domain.SectionScaffold{Content: "new section 2\n"})
	assert.Equal(t, domain.SectionCollection{sections[0], newSection2}, sections.Replace(sections[1], newSection2))
}

func TestSectionCollectionText(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	assert.Equal(t, "section 1\nsection 2\n", sections.Text())
}
