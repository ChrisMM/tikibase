package domain

import (
	"fmt"
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
		data.Doc = ScaffoldDocument(DocumentScaffold{})
	}
	return Section{content: SectionContent(data.Content), document: data.Doc}
}

// Anchor provides the URL anchor for this TikiSection.
func (s *Section) Anchor() (string, error) {
	sectionTitle, err := s.Title()
	if err != nil {
		return "", errors.Wrap(err, "cannot determine the section anchor")
	}
	return strcase.ToKebab(sectionTitle), nil
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
	sectionTitle, err := s.Title()
	if err != nil {
		return result, err
	}
	matches := markdownLinkRE.FindAllStringSubmatch(string(s.content), 9999)
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
		targetFileName := DocumentFilename(filename)
		targetDocument, err := tdc.Find(targetFileName)
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s' in Section '%s'", targetFileName, linkTitle, sectionTitle)
		}
		result = append(result, newTikiLink(linkTitle, s, targetDocument))
	}
	matches = htmlLinkRE.FindAllStringSubmatch(string(s.content), 9999)
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
		targetDocument, err := tdc.Find(DocumentFilename(targetFilename))
		if err != nil {
			return result, errors.Wrapf(err, "cannot find target document ('%s') for link '%s'", targetFilename, linkTitle)
		}
		result = append(result, newTikiLink(linkTitle, s, targetDocument))
	}

	return result.Unique(), nil
}

// Title returns the human-friendly title of this TikiSection,
// i.e. its title tag without the "###"" in front
func (s *Section) Title() (string, error) {
	titleLine := strings.SplitN(string(s.content), "\n", 1)[0]
	matches := stripTitleTagRE.FindStringSubmatch(titleLine)
	if len(matches) == 0 {
		return "", fmt.Errorf("malformatted section title: '%s'", titleLine)
	}
	return matches[1], nil
}

// URL provides the full URL to this Section (link to document that contains this section + anchor of the section.
func (s *Section) URL() (string, error) {
	sectionAnchor, err := s.Anchor()
	if err != nil {
		return "", errors.Wrap(err, "cannot determine the URL for section")
	}
	return s.document.URL() + "#" + sectionAnchor, nil
}
