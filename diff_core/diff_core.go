package diff_core

func CompareLine(line1 string, line2 string) CompareLineResult {
	changes := []Change{}
	line1Runes := []rune(line1)
	line2Runes := []rune(line2)

	for pos, char := range line1Runes {
		altChar := line2Runes[pos]
		if char != altChar {
			changes = append(changes, Change{pos, pos})
		}
	}

	return CompareLineResult{line1 == line2, changes}
}

type CompareLineResult struct {
	IsSame  bool
	Changes []Change
}

type Change struct {
	Start int
	End   int
}
