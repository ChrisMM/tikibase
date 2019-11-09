package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kevgo/tikibase/domain"
)

func TestSectionCollection_FindByTitle(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "# Title\ntext 1\n"},
		{Content: "### what is it\ntext 2\n"},
		{Content: "### what does it\ntext 3\n"},
	})

	actual, err := sections.FindByTitle("what is it")

	assert.Nil(t, err)
	assert.Same(t, sections[1], actual)
}

func TestSectionCollection_Remove(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{{}, {}})
	assert.Equal(t, domain.SectionCollection{sections[0]}, sections.Remove(sections[1]))
}

func TestSectionCollection_Replace(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{{}, {}})
	newSection2 := domain.ScaffoldSection(domain.SectionScaffold{Content: "new section 2\n"})
	assert.Equal(t, domain.SectionCollection{sections[0], newSection2}, sections.Replace(sections[1], newSection2))
}

func TestSectionCollection_Text(t *testing.T) {
	sections := domain.ScaffoldSectionCollection([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	assert.Equal(t, "section 1\nsection 2\n", sections.Text())
}
