package linkify

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
	"github.com/stretchr/testify/assert"
)

func TestLinkify_noMatch(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "aws.md", Content: "# Amazon Web Services\n"},
		{FileName: "text.md", Content: "# Text\n\nthe text"},
	})
	mappings := []docMapping{{lookFor: "Amazon Web Services", replaceWith: "[Amazon Web Services](aws.md)"}}
	have, err := linkifyDoc(docs[1], docs, mappings)
	assert.Nil(t, err)
	assert.Equal(t, "# Text\n\nthe text", have)
}

func TestLinkify_match(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "aws.md", Content: "# Amazon Web Services\n"},
		{FileName: "text.md", Content: "# Text\n\nAmazon Web Services is a cloud provider"},
	})
	mappings := []docMapping{{lookFor: "Amazon Web Services", replaceWith: "[Amazon Web Services](aws.md)"}}
	have, err := linkifyDoc(docs[1], docs, mappings)
	assert.Nil(t, err)
	assert.Equal(t, "# Text\n\n[Amazon Web Services](aws.md) is a cloud provider", have)
}

func TestLinkify_match_differentCase(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "aws.md", Content: "# Amazon Web Services\n"},
		{FileName: "text.md", Content: "# Text\n\nAmazon web services is a cloud provider"},
	})
	mappings := []docMapping{{lookFor: "Amazon Web Services", replaceWith: "[Amazon Web Services](aws.md)"}}
	have, err := linkifyDoc(docs[1], docs, mappings)
	assert.Nil(t, err)
	assert.Equal(t, "# Text\n\n[Amazon Web Services](aws.md) is a cloud provider", have)
}

func TestLinkify_matchWithExistingLinks(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "aws.md", Content: "# Amazon Web Services\n"},
		{FileName: "text.md", Content: "# Text\n\n[Amazon Web Services](aws.md) is a cloud provider. Amazon Web Services is also the biggest product line of Amazon."},
	})
	mappings := []docMapping{{lookFor: "Amazon Web Services", replaceWith: "[Amazon Web Services](aws.md)"}}
	have, err := linkifyDoc(docs[1], docs, mappings)
	assert.Nil(t, err)
	assert.Equal(t, "# Text\n\n[Amazon Web Services](aws.md) is a cloud provider. Amazon Web Services is also the biggest product line of Amazon.", have)
}

func TestLinkify_multiMatchWithSubMatch(t *testing.T) {
	docs := domain.ScaffoldDocuments([]domain.DocumentScaffold{
		{FileName: "security-user-stories.md", Content: "# Security User Stories\n"},
		{FileName: "user-stories.md", Content: "# User Stories\n"},
		{FileName: "text.md", Content: "# Text\n\nlink: [security user stories](security-user-stories.md), no link: security user stories"},
	})
	mappings := []docMapping{
		{lookFor: "security user stories", replaceWith: "[security-user-stories](security-user-stories.md)"},
		{lookFor: "user stories", replaceWith: "[user-stories](user-stories.md)"},
	}
	have, err := linkifyDoc(docs[2], docs, mappings)
	assert.Nil(t, err)
	assert.Equal(t, "# Text\n\nlink: [security user stories](security-user-stories.md), no link: security user stories", have)
}
