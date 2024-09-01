package diff_core_test

import (
	core "diff_cli/diff_core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareLinesSingleDifference(t *testing.T) {
	result := core.CompareLine("a", "b")
	assert.False(t, result.IsSame)
}

func TestSingleCharacterDifference(t *testing.T) {
	result := core.CompareLine("abc", "adc")

	assert.Equal(t, []core.Change{{1, 1}}, result.Changes)
}

func TestTwoCharacterDifference(t *testing.T) {
	result := core.CompareLine("abcd", "azzd")

	assert.Equal(t, []core.Change{{1, 2}}, result.Changes)
}

func TestTwoLotsOfDifferences(t *testing.T) {
	result := core.CompareLine("abcdefgh", "azzdezzh")

	assert.Equal(t, []core.Change{{1, 2}, {5, 6}}, result.Changes)
}

func TestCharacterDifferencesAtTheEndOfTheLine(t *testing.T) {
	result := core.CompareLine("abc", "azz")

	assert.Equal(t, []core.Change{{1, 2}}, result.Changes)
}

func TestWhenFirstLineIsShorter(t *testing.T) {
	result := core.CompareLine("abc", "abcd")

	// Start Index past the last element of line1
	// End Index is the last element of line 2
	assert.Equal(t, []core.Change{{3, 3}}, result.Changes)
}

func TestWhenSecondLineIsShorter(t *testing.T) {
	result := core.CompareLine("abcd", "abc")

	// Start Index past the last element of line1
	// End Index is the last element of line 2
	assert.Equal(t, []core.Change{{3, 3}}, result.Changes)
}
