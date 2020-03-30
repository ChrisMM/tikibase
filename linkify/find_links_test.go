package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLinks_multipleLinks(t *testing.T) {
	text := `[Amazon Web Services](aws.md) is a cloud provider.
					 [Google Cloud](gcp.md) is another one.
					 But [Amazon Web Services](aws.md) is bigger.`
	have := findLinks(text)
	want := []string{"[Amazon Web Services](aws.md)", "[Google Cloud](gcp.md)"}
	assert.Equal(t, want, have)
}

func TestFindLinks_sorting(t *testing.T) {
	text := `[Amazon](amazon.md) makes [Amazon Web Services](aws.md)`
	have := findLinks(text)
	want := []string{"[Amazon Web Services](aws.md)", "[Amazon](amazon.md)"}
	assert.Equal(t, want, have)
}
