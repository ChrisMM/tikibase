package config_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/kevgo/tikibase/config"
	"github.com/stretchr/testify/assert"
)

func TestIgnores(t *testing.T) {
	tests := map[string]bool{
		"Makefile":     true,
		"tikibase.yml": true,
		".gitignore":   true,
		"foo.md":       false,
		"img.png":      false,
	}
	conf := config.Scaffold([]string{"Makefile"})
	for give := range tests {
		ignores, err := conf.Ignores(give)
		assert.Nil(t, err)
		assert.Equal(t, tests[give], ignores)
	}
}

func TestLoad(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	assert.Nil(t, err)
	err = ioutil.WriteFile(filepath.Join(dir, config.FileName()), []byte("ignore:\n  - Makefile\n"), 0644)
	assert.Nil(t, err)
	conf, err := config.Load(dir)
	assert.Nil(t, err)
	ignores, err := conf.Ignores("Makefile")
	assert.Nil(t, err)
	assert.True(t, ignores)
}
