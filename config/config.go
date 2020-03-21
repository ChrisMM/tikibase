package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Schema defines the data schema of the configuration.
type Schema struct {
	Ignore []string // must be exported for the YML parser
}

// defaults provides the default configuration values to be used in the absence of user-provided configuration.
func defaults() Schema {
	return Schema{
		Ignore: []string{"node_modules", "vendor"},
	}
}

// FileName provides the name of the config file.
func FileName() string {
	return "tikibase.yml"
}

// Scaffold provides configuration instances for testing.
func Scaffold(ignore []string) Schema {
	return Schema{ignore}
}

// Ignores indicates whether this configuration
// is set to ignore the given filename.
func (schema Schema) Ignores(filename string) (result bool, err error) {
	// ignore hidden files
	if strings.HasPrefix(filename, ".") {
		return true, nil
	}
	// we always ignore the config file itself
	if filename == FileName() {
		return true, nil
	}
	for _, ignore := range schema.Ignore {
		match, err := filepath.Match(ignore, filename)
		if err != nil {
			return result, fmt.Errorf("cannot match filename %q to %q: %w", filename, ignore, err)
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}

// Load provides the configuration values
// for the TikiBase in the given directory.
func Load(dir string) (result Schema, err error) {
	data, err := ioutil.ReadFile(filepath.Join(dir, FileName()))
	if err != nil {
		return defaults(), nil
	}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("invalid config file content: %w", err)
	}
	return result, nil
}
