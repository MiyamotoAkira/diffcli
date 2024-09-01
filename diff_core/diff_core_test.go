package diff_core_test

import (
	core "diff_cli/diff_core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareLinesSingleDifference(t *testing.T) {
	result := core.CompareLine("a", "b")
	assert.False(t, result)
}
