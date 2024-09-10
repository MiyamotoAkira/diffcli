package diff_cli

import (
	"bufio"
	"fmt"
	core "github.com/MiyamotoAkira/diffcli/diff_core"
	"os"
	"strings"
)

func CompareFiles(file1Name string, file2Name string) string {

	file1, err := os.Open(file1Name)
	if err != nil {
		return "Error"
	}
	defer file1.Close()
	file1Scanner := bufio.NewScanner(file1)
	file2, err := os.Open(file2Name)
	if err != nil {
		return "Error"
	}
	defer file2.Close()
	file2Scanner := bufio.NewScanner(file2)

	var file1Lines []string
	for file1Scanner.Scan() {
		file1Lines = append(file1Lines, file1Scanner.Text())
	}
	var file2Lines []string
	for file2Scanner.Scan() {
		file2Lines = append(file2Lines, file2Scanner.Text())
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
