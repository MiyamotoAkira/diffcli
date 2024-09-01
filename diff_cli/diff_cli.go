package diff_cli

import (
	"fmt"
	"os"
	"strings"

	core "github.com/MiyamotoAkira/diffcli/diff_core"
)

func CompareFiles(file1Name string, file2Name string) string {

	file1Content, err := os.ReadFile(file1Name)

	if err != nil {
		return "Error"
	}

	file2Content, err := os.ReadFile(file2Name)

	if err != nil {
		return "Error"
	}

	file1Lines := strings.Split(string(file1Content), "\n")
	file2Lines := strings.Split(string(file2Content), "\n")
	println(string(file1Content))
	println(file1Lines)
	println("********************")
	println(len(file1Lines))
	println("********************")
	println(string(file2Content))
	println(file2Lines)
	println("********************")
	println(len(file2Lines))
	println("********************")

	result := core.CompareLines(file1Lines, file2Lines)

	// {[]core.Change{}},
	// {[]core.Change{{1, 1}}},
	// {[]core.Change{{0, 2}}},

	println("********************")
	println(len(result))
	println("********************")
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
			output.WriteString("\n")
		}
	}

	return output.String()
}
