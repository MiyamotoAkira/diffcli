package diff_io

import (
	"bufio"
	"fmt"
	core "github.com/MiyamotoAkira/diffcli/diff_core"
	"os"
	"strings"
)

func CompareFiles(file1Name string, file2Name string) string {
	file1Lines, err1 := readFile(file1Name)
	file2Lines, err2 := readFile(file2Name)
	if err1 != nil {
		return "error cannot read file 1 " + err1.Error()
	}
	if err2 != nil {
		return "error cannot read file 2 " + err2.Error()
	}

	result := core.CompareLines(file1Lines, file2Lines)

	var output strings.Builder
	for pos, change := range result {
		if !change.IsSame() {
			if pos < len(file1Lines) {
				output.WriteString(fmt.Sprintf("- %s", file1Lines[pos]))
			} else {
				output.WriteString("-")
			}
			output.WriteString("\n")

			if pos < len(file2Lines) {
				output.WriteString(fmt.Sprintf("+ %s", file2Lines[pos]))
			} else {
				output.WriteString("+")
			}
			if pos < len(result)-1 {
				output.WriteString("\n")
			}
		}
	}

	return output.String()
}

func readFile(file1Name string) ([]string, error) {
	file1, err := os.Open(file1Name)
	if err != nil {
		return nil, err
	}
	defer file1.Close()
	file1Scanner := bufio.NewScanner(file1)
	var file1Lines []string
	for file1Scanner.Scan() {
		file1Lines = append(file1Lines, file1Scanner.Text())
	}
	return file1Lines, nil
}

func CompareDirectories(directory1 string, directory2 string) string {
	filesDirectory1, err := os.ReadDir(directory1)
	if err != nil {
		return "Failed reading " + directory1
	}

	filesDirectory2, err := os.ReadDir(directory2)
	if err != nil {
		return "Failed reading " + directory2
	}

	var filesDirectory1Map = make(map[string]string)
	var filesDirectory2Map = make(map[string]string)

	for _, entry := range filesDirectory1 {
		filesDirectory1Map[entry.Name()] = entry.Name()
	}

	for _, entry := range filesDirectory2 {
		filesDirectory2Map[entry.Name()] = entry.Name()
	}

	var onlyInDirectory1 []string
	var onlyInDirectory2 []string

	for k := range filesDirectory1Map {
		_, ok := filesDirectory2Map[k]
		if !ok {
			onlyInDirectory1 = append(onlyInDirectory1, k)
		}
	}

	for k := range filesDirectory2Map {
		_, ok := filesDirectory1Map[k]
		if !ok {
			onlyInDirectory2 = append(onlyInDirectory2, k)
		}
	}

	var output strings.Builder

	for _, entry := range onlyInDirectory1 {
		output.WriteString(fmt.Sprintf("Only in %s: %s\n", directory1, entry))
	}

	for _, entry := range onlyInDirectory2 {
		output.WriteString(fmt.Sprintf("Only in %s: %s\n", directory2, entry))
	}

	return output.String()
}
