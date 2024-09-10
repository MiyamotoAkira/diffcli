package diff_cli_test

import (
	cli "github.com/MiyamotoAkira/diffcli/diff_cli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareTwoFiles(t *testing.T) {
	var file1Name = "file1.txt"
	var file2Name = "file2.txt"
	var result string

	result = cli.CompareFiles(file1Name, file2Name)
	assert.Equal(t, "- def\n+ dzf\n- ghi\n+", result)
}
