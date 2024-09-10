package diff_cli_test

import (
	cli "github.com/MiyamotoAkira/diffcli/diff_cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type FileTestSuite struct {
	suite.Suite
	folderName string
}

func (suite *FileTestSuite) SetupSuite() {
	folder, err := os.MkdirTemp("./", "temp")
	if err != nil {
		panic("Error creating directory")
	}
	suite.folderName = folder
}

func (suite *FileTestSuite) TearDownSuite() {
	err := os.RemoveAll(suite.folderName)

	if err != nil {
		panic("Issue removing folder " + suite.folderName)
	}
}

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

func (suite *FileTestSuite) TestCompareTwoFiles() {
	var file1Name = suite.folderName + "/" + "file1.txt"
	var file2Name = suite.folderName + "/" + "file2.txt"
	createFile(file1Name, []string{"abc", "def", "ghi"})
	createFile(file2Name, []string{"abc", "dzf"})
	var result string

	result = cli.CompareFiles(file1Name, file2Name)
	assert.Equal(suite.T(), "- def\n+ dzf\n- ghi\n+", result)

}

func TestFileTestSuite(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}
