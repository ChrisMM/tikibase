package check_test

import (
	"testing"

	"github.com/kevgo/tikibase/check"
	"github.com/stretchr/testify/assert"
)

func TestFindGroups_hasClusters(t *testing.T) {
	groups := check.FindGroups([]string{"one", "One", "ONE", "two", "Two"})
	expected := [][]string{{"one", "One", "ONE"}, {"two", "Two"}}
	assert.ElementsMatch(t, groups, expected)
}

func TestFindGroups_noClusters(t *testing.T) {
	groups := check.FindGroups([]string{"one", "two", "three"})
	assert.ElementsMatch(t, groups, [][]string{})
}
