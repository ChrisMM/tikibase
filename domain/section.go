package domain

import (
	"regexp"
	"strings"

	"github.com/kevgo/tikibase/helpers"

	"github.com/iancoleman/strcase"

	"github.com/pkg/errors"
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
func ScaffoldSection(data SectionScaffold) Section {
	if data.Content == "" {
		data.Content = "### default section\ncontent\n"
	}
	if data.Doc == nil {
		doc := ScaffoldDocument(DocumentScaffold{})
		data.Doc = &doc
	}
	return Section{content: SectionContent(data.Content), document: data.Doc}
}

// Anchor provides the URL anchor for this TikiSection.
func (s *Section) Anchor() string {
	return strcase.ToKebab(s.Title())
}

// AppendLine provides a new Section that is this Section with the given line appended.
func (s *Section) AppendLine(line string) Section {
	return Section{content: s.content + SectionContent(line), document: s.document}
}

// Content returns the complete content of the entire section.
func (s *Section) Content() SectionContent {
	return s.content
}

// Document provides a link to the Document containing this section.
func (s *Section) Document() *Document {
	return s.document
}

// TikiLinks returns all TikiLinks in this section.
func (s *Section) TikiLinks(tdc DocumentCollection) (result TikiLinkCollection, err error) {
	matches := markdownLinkRE.FindAllStringSubmatch(string(s.content), 9999)
	for i := range matches {
		linkTitle := matches[i][1]
		linkTarget := matches[i][2]
		if helpers.IsURL(linkTarget) {
			continue
		}
		filename, _ := helpers.SplitURL(linkTarget)
		targetFileName := DocumentFilename(filename)
		targetDocument, err := tdc.Find(targetFileName)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s' in Section '%s'", targetFileName, linkTitle, s.Title())
		}
		result = append(result, newTikiLink(linkTitle, s, targetDocument))
	}
	matches = htmlLinkRE.FindAllStringSubmatch(string(s.content), 9999)
	for _, match := range matches {
		linkTitle := match[2]
		targetFilename := DocumentFilename(match[1])
		targetDocument, err := tdc.Find(targetFilename)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetFilename, linkTitle)
		}
		result = append(result, newTikiLink(linkTitle, s, targetDocument))
	}
	return result, nil
}

// Title returns the human-friendly title of this TikiSection,
// i.e. its title tag without the "###"" in front
func (s *Section) Title() string {
	titleLine := strings.SplitN(string(s.content), "\n", 1)[0]
	matches := stripTitleTagRE.FindStringSubmatch(titleLine)
	return matches[1]
}

// URL provides the full URL to this Section (link to document that contains this section + anchor of the section.
func (s *Section) URL() string {
	return s.document.URL() + "#" + s.Anchor()
}
