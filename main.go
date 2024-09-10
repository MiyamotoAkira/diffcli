package main

import (
	io "github.com/MiyamotoAkira/diffcli/diff_io"
	"os"
)

func main() {
	filePath1 := os.Args[1]
	filePath2 := os.Args[2]

	result := io.CompareFiles(filePath1, filePath2)

	println(result)
}
