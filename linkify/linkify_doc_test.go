package linkify

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestLinkify_noMatch(t *testing.T) {
	mappings := []docMapping{{name: "the link", file: "target"}}
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{Content: "the text"})
	have := linkifyDoc(doc, mappings)
	assert.Equal(t, "the text", have)
}

func TestLinkify_match(t *testing.T) {
	mappings := []docMapping{{name: "Amazon Web Services", file: "aws.md"}}
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{Content: "Amazon Web Services is a cloud provider"})
	have := linkifyDoc(doc, mappings)
	assert.Equal(t, "[Amazon Web Services](aws.md) is a cloud provider", have)
}

func TestLinkify_match_differentCase(t *testing.T) {
	mappings := []docMapping{{name: "Amazon Web Services", file: "aws.md"}}
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{Content: "Amazon web services is a cloud provider"})
	have := linkifyDoc(doc, mappings)
	assert.Equal(t, "[Amazon Web Services](aws.md) is a cloud provider", have)
}

func TestLinkify_matchWithExistingLinks(t *testing.T) {
	mappings := []docMapping{{name: "Amazon Web Services", file: "aws.md"}}
	doc := domain.ScaffoldDocument(domain.DocumentScaffold{Content: `[Amazon Web Services](aws.md) is a cloud provider. Amazon Web Services is also the biggest product line of Amazon.`})
	have := linkifyDoc(doc, mappings)
	assert.Equal(t, "[Amazon Web Services](aws.md) is a cloud provider. [Amazon Web Services](aws.md) is also the biggest product line of Amazon.", have)
}
