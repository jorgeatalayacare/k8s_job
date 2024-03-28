package main

import (
	"fmt"
	"os"

	fr "github.com/jorgeatalayacare/k8s_job/file_reader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file path>\n", os.Args[0])
		os.Exit(1)
	}
	filePath := os.Args[1]

	fr.ReadPrintFile(filePath)
}
