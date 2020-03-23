package linkify_test

import (
	"testing"

	"github.com/kevgo/tikibase/linkify"
	"github.com/stretchr/testify/assert"
)

func TestLinkify_noMatch(t *testing.T) {
	have := linkify.Linkify("the text", "the link", "target")
	assert.Equal(t, "the text", have)
}

func TestLinkify_match(t *testing.T) {
	have := linkify.Linkify("Amazon Web Services is a cloud provider", "Amazon Web Services", "aws.md")
	assert.Equal(t, "[Amazon Web Services](aws.md) is a cloud provider", have)
}

func TestLinkify_match_differentCase(t *testing.T) {
	have := linkify.Linkify("Amazon web services is a cloud provider", "Amazon Web Services", "aws.md")
	assert.Equal(t, "[Amazon Web Services](aws.md) is a cloud provider", have)
}

func TestLinkify_matchWithExistingLinks(t *testing.T) {
	text := `[Amazon Web Services](aws.md) is a cloud provider. Amazon Web Services is also the biggest product line of Amazon.`
	have := linkify.Linkify(text, "Amazon Web Services", "aws.md")
	assert.Equal(t, "[Amazon Web Services](aws.md) is a cloud provider. [Amazon Web Services](aws.md) is also the biggest product line of Amazon.", have)
}
