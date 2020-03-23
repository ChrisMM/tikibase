package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestSections_FindByTitle(t *testing.T) {
	sections := domain.ScaffoldSections([]domain.SectionScaffold{
		{Content: "# Title\ntext 1\n"},
		{Content: "### what is it\ntext 2\n"},
		{Content: "### what does it\ntext 3\n"},
	})
	actual, err := sections.FindByTitle("what is it")
	assert.Nil(t, err)
	assert.Same(t, sections[1], actual)
}

func TestSections_Remove(t *testing.T) {
	sections := domain.ScaffoldSections([]domain.SectionScaffold{{}, {}})
	assert.Equal(t, domain.Sections{sections[0]}, sections.Remove(sections[1]))
}

func TestSections_Replace(t *testing.T) {
	sections := domain.ScaffoldSections([]domain.SectionScaffold{{}, {}})
	newSection2 := domain.ScaffoldSection(domain.SectionScaffold{Content: "new section 2\n"})
	assert.Equal(t, domain.Sections{sections[0], newSection2}, sections.Replace(sections[1], newSection2))
}

func TestSections_Text(t *testing.T) {
	sections := domain.ScaffoldSections([]domain.SectionScaffold{
		{Content: "section 1\n"},
		{Content: "section 2\n"},
	})
	assert.Equal(t, "section 1\nsection 2\n", sections.Text())
}
