package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLinks(t *testing.T) {
	text := `[Amazon Web Services](aws.md) is a cloud provider.
					 [Google Cloud](gcp.md) is another one.
					 But [Amazon Web Services](aws.md) is bigger.`
	have := findLinks(text)
	assert.Equal(t, []string{"[Amazon Web Services](aws.md)", "[Google Cloud](gcp.md)"}, have)
}
