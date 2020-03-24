package linkify_test

import (
	"testing"

	"github.com/kevgo/tikibase/linkify"
	"github.com/stretchr/testify/assert"
)

func TestFindExistingLinks(t *testing.T) {
	text := `[Amazon Web Services](aws.md) is a cloud provider.
					 [Google Cloud](gcp.md) is another one.
					 But [Amazon Web Services](aws.md) is bigger.`
	have := linkify.FindExistingLinks(text)
	assert.Equal(t, []string{"[Amazon Web Services](aws.md)", "[Google Cloud](gcp.md)"}, have)
}
