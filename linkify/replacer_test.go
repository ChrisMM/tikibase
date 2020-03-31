package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacer(t *testing.T) {
	ir := NewIgnoringReplacer()
	ir.Ignore("# Amazon\n")
	ir.ReplaceOnce("Amazon", "[Amazon](amazon.md)")
	ir.ReplaceOnce("Amazon Web Services", "[Amazon Web Services](aws.md)")
	give := "# Amazon\n\nAmazon web services is a product of Amazon. Amazon web services was invented at Amazon."
	have := ir.Apply(give)
	want := "# Amazon\n\n[Amazon Web Services](aws.md) is a product of [Amazon](amazon.md). Amazon web services was invented at Amazon."
	assert.Equal(t, want, have)
}

func TestReplarer_Ignores(t *testing.T) {
	ir := NewIgnoringReplacer()
	ir.Ignore("foo")
	assert.True(t, ir.Ignores("foo"))
	assert.False(t, ir.Ignores("bar"))
}
