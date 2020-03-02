package domain

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kevgo/tikibase/helpers"
)

// Section represents a section in a Document.
type Section struct {
	// the textual content of this TikiSection
	// including the title line
	content SectionContent

	// document links to the Document that contains this Section.
	document *Document
}

// SectionScaffold defines the named arguments for ScaffoldSection.
type SectionScaffold struct {
	Content string
	Doc     *Document
}

// SectionContent represents the full content of a TikiSection,
// including heading tag and body.
type SectionContent string

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var markdownLinkRE = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var htmlLinkRE = regexp.MustCompile(`<a[^>]* href="(.*?)"[^>]*>(.*?)</a>`)

// This global variable is a constant and doesn't need to be stubbed in tests.
//nolint:gochecknoglobals
var stripTitleTagRE = regexp.MustCompile(`#+\s*(.*)`)

// ScaffoldSection provides new TikiSection instances for testing.
func ScaffoldSection(data SectionScaffold) *Section {
	if data.Content == "" {
		data.Content = "### default section\ncontent\n"
	}
	if data.Doc == nil {
		data.Doc = ScaffoldDocument(DocumentScaffold{})
	}
	return &Section{content: SectionContent(data.Content), document: data.Doc}
}

// Anchor provides the URL anchor for this TikiSection.
func (section *Section) Anchor() (string, error) {
	sectionTitle, err := section.Title()
	if err != nil {
		return "", fmt.Errorf("cannot determine the section anchor: %w", err)
	}
	return strcase.ToKebab(sectionTitle), nil
}

// AppendText provides a new Section that is this Section with the given line appended.
func (section *Section) AppendText(line string) Section {
	return Section{content: section.content + SectionContent(line), document: section.document}
}

// Content returns the complete content of the entire section.
func (section *Section) Content() SectionContent {
	return section.content
}

// Document provides a link to the Document containing this section.
func (section *Section) Document() *Document {
	return section.document
}

// TikiLinks returns all TikiLinks in this section.
func (section *Section) TikiLinks(dc DocumentCollection) (result TikiLinkCollection, err error) {
	sectionTitle, err := section.Title()
	if err != nil {
		return result, err
	}
	matches := markdownLinkRE.FindAllStringSubmatch(string(section.content), 9999)
	for i := range matches {
		linkTitle := matches[i][1]
		linkTarget := matches[i][2]
		if helpers.IsURL(linkTarget) {
			// we can ignore links to external files here
			continue
		}
		if strings.HasPrefix(linkTarget, "#") {
			// we can ignore links within the same file here
			continue
		}
		filename, _ := helpers.SplitURL(linkTarget)
		if !strings.HasSuffix(filename, ".md") {
			// we can ignore links to non-Markdown files here
			continue
		}
		targetDocument, err := dc.FindByFilename(DocumentFilename(filename))
		if err != nil {
			return result, fmt.Errorf("cannot find target document (%q) for link %q in Section %q: %w", filename, linkTitle, sectionTitle, err)
		}
		result = append(result, newTikiLink(linkTitle, section, targetDocument))
	}
	matches = htmlLinkRE.FindAllStringSubmatch(string(section.content), 9999)
	for _, match := range matches {
		linkTitle := match[2]
		targetFilename := match[1]
		if helpers.IsURL(targetFilename) {
			// we can ignore links to external files here
			continue
		}
		if strings.HasPrefix(targetFilename, "#") {
			// we can ignore links within the same file here
			continue
		}
		if !strings.HasSuffix(targetFilename, ".md") {
			// we can ignore links to non-Markdown files here
			continue
		}
		targetDocument, err := dc.FindByFilename(DocumentFilename(targetFilename))
		if err != nil {
			return result, fmt.Errorf("cannot find target document (%q) for link %q: %w", targetFilename, linkTitle, err)
		}
		result = append(result, newTikiLink(linkTitle, section, targetDocument))
	}

	return result.Unique(), nil
}

// Title returns the human-friendly title of this TikiSection,
// i.e. its title tag without the "###"" in front
func (section *Section) Title() (string, error) {
	titleLine := strings.SplitN(string(section.content), "\n", 1)[0]
	matches := stripTitleTagRE.FindStringSubmatch(titleLine)
	if len(matches) == 0 {
		return "", fmt.Errorf("malformatted section title: %q", titleLine)
	}
	return matches[1], nil
}

// URL provides the full URL to this Section (link to document that contains this section + anchor of the section.
func (section *Section) URL() (string, error) {
	sectionAnchor, err := section.Anchor()
	if err != nil {
		return "", fmt.Errorf("cannot determine the URL for section: %w", err)
	}
	return section.document.URL() + "#" + sectionAnchor, nil
}
