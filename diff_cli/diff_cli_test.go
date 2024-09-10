package diff_cli_test

import (
	cli "github.com/MiyamotoAkira/diffcli/diff_cli"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func createFile(fileName string, fileContent []string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic("Creating file " + fileName)
	}
	defer f.Close()

	for _, line := range fileContent {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			panic("Writing file " + fileName)
		}
	}
}

func TestCompareTwoFiles(t *testing.T) {
	folder, err := os.MkdirTemp("./", "temp")
	if err != nil {
		panic("Error creating directory")
	}
	var file1Name = folder + "/" + "file1.txt"
	var file2Name = folder + "/" + "file2.txt"
	createFile(file1Name, []string{"abc", "def", "ghi"})
	createFile(file2Name, []string{"abc", "dzf"})
	var result string

	result = cli.CompareFiles(file1Name, file2Name)
	assert.Equal(t, "- def\n+ dzf\n- ghi\n+", result)

	err = os.RemoveAll(folder)

	if err != nil {
		panic("Issue removing folder " + folder)
	}
}
