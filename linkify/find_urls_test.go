package linkify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUrls(t *testing.T) {
	text := "google: http://google.com, amazon: https://amazon.com"
	have := findUrls(text)
	want := []string{
		"http://google.com",
		"https://amazon.com",
	}
	assert.Equal(t, want, have)
}
