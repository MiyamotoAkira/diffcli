package diff_core

func CompareLines(lines1 []string, lines2 []string) []CompareLineResult {
	results := []CompareLineResult{}

	for pos, line := range lines1 {
		if pos > len(lines2)-1 {
			for i := range len(lines1) - len(lines2) {
				index := len(lines2) + i
				results = append(results, CompareLineResult{[]Change{{0, len(lines1[index]) - 1}}})
			}
			break
		}
		results = append(results, CompareLine(line, lines2[pos]))
	}

	if len(lines2) > len(lines1) {
		for i := range len(lines2) - len(lines1) {
			index := len(lines1) + i
			results = append(results, CompareLineResult{[]Change{{0, len(lines2[index]) - 1}}})
		}
	}

	return results
}

func CompareLine(line1 string, line2 string) CompareLineResult {
	changes := []Change{}
	line1Runes := []rune(line1)
	line2Runes := []rune(line2)

	var onChanges bool
	var startIndex int

	for pos, char := range line1Runes {

		if pos > len(line2Runes)-1 {
			if !onChanges {
				onChanges = true
				startIndex = pos
			}
			break
		}

		altChar := line2Runes[pos]
		if char != altChar {
			if !onChanges {
				onChanges = true
				startIndex = pos
			}
		} else {
			if onChanges {
				onChanges = false
				changes = append(changes, Change{startIndex, pos - 1})
			}
		}

	}
	if onChanges {
		onChanges = false
		changes = append(changes, Change{startIndex, len(line1Runes) - 1})
	}

	if len(line2Runes) > len(line1Runes) {
		changes = append(changes, Change{len(line1Runes), len(line2Runes) - 1})
	}

	return CompareLineResult{changes}
}

type CompareLineResult struct {
	Changes []Change
}

func (result *CompareLineResult) IsSame() bool {
	return len(result.Changes) == 0
}

type Change struct {
	Start int
	End   int
}
