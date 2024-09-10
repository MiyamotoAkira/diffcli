package diff_io_test

import (
	"fmt"
	cli "github.com/MiyamotoAkira/diffcli/diff_io"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type FileTestSuite struct {
	suite.Suite
	folderName string
	file1Name  string
	file2Name  string
}

func (suite *FileTestSuite) SetupTest() {
	folder, err := os.MkdirTemp("./", "temp")
	if err != nil {
		panic("Error creating directory")
	}
	suite.folderName = folder
	suite.file1Name = suite.folderName + "/" + "file1.txt"
	suite.file2Name = suite.folderName + "/" + "file2.txt"
}

func (suite *FileTestSuite) TearDownTest() {
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
	createFile(suite.file1Name, []string{"abc", "def", "ghi"})
	createFile(suite.file2Name, []string{"abc", "dzf"})

	result := cli.CompareFiles(suite.file1Name, suite.file2Name)

	suite.Equal("- def\n+ dzf\n- ghi\n+", result)
}

func (suite *FileTestSuite) TestCompareEmptyFiles() {
	createFile(suite.file1Name, []string{""})
	createFile(suite.file2Name, []string{""})

	result := cli.CompareFiles(suite.file1Name, suite.file2Name)

	suite.Equal("", result)
}

func (suite *FileTestSuite) TestCompareEqualFiles() {
	createFile(suite.file1Name, []string{"darth vader", "luke skywalker", "han solo"})
	createFile(suite.file2Name, []string{"darth vader", "luke skywalker", "han solo"})

	result := cli.CompareFiles(suite.file1Name, suite.file2Name)

	suite.Equal("", result)
}

func (suite *FileTestSuite) TestSecondFileDoesNotExist() {
	createFile(suite.file1Name, []string{"abc", "def", "ghi"})
	//do not create second file

	result := cli.CompareFiles(suite.file1Name, suite.file2Name)

	suite.Contains(result, "error cannot read file 2")
}

func (suite *FileTestSuite) TestFirstFileDoesNotExist() {
	//do not create first file
	createFile(suite.file2Name, []string{"abc", "def", "ghi"})

	result := cli.CompareFiles(suite.file1Name, suite.file2Name)

	suite.Contains(result, "error cannot read file 1")
}

func TestFileTestSuite(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}

type DirectoryTestSuite struct {
	suite.Suite
	folder1Name string
	folder2Name string
}

func (suite *DirectoryTestSuite) SetupTest() {
	folder, err := os.MkdirTemp("./", "temp")
	if err != nil {
		panic("Error creating directory")
	}
	suite.folder1Name = folder

	folder, err = os.MkdirTemp("./", "temp")
	if err != nil {
		panic("Error creating directory")
	}
	suite.folder2Name = folder
}

func (suite *DirectoryTestSuite) TearDownTest() {
	err := os.RemoveAll(suite.folder1Name)

	if err != nil {
		panic("Issue removing folder " + suite.folder1Name)
	}

	err = os.RemoveAll(suite.folder2Name)

	if err != nil {
		panic("Issue removing folder " + suite.folder2Name)
	}
}

func (suite *DirectoryTestSuite) TestTwoEmptyDirectory() {

	result := cli.CompareDirectories(suite.folder1Name, suite.folder2Name)

	suite.Equal("", result)
}

func (suite *DirectoryTestSuite) TestCompareCompletelyDifferentDirectories() {
	file1Name := "file1.test"
	file2Name := "file2.test"
	file1Path := suite.folder1Name + "/" + file1Name
	file2Path := suite.folder2Name + "/" + file2Name

	createFile(file1Path, []string{})
	createFile(file2Path, []string{})
	result := cli.CompareDirectories(suite.folder1Name, suite.folder2Name)

	// Probably the last \n shouldn't be there
	suite.Equal(fmt.Sprintf("Only in %s: %s\nOnly in %s: %s\n", suite.folder1Name, file1Name, suite.folder2Name, file2Name), result)
}

func TestDirectoryTestSuite(t *testing.T) {
	suite.Run(t, new(DirectoryTestSuite))
}
