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
	if err1 != nil || err2 != nil {
		return "error"
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
