package external

import (
	"bufio"
	"os"
	"strings"
)

func SortByMergeWithPresort(inputFilename, outputFilename string, n int) {
	tempFilenames := []string{"temp-0.text", "temp-1.text"}
	presort(inputFilename, tempFilenames[0], n, 100)

	for l := 100; l < n; l *= 2 {
		singleMergePass(tempFilenames[0], tempFilenames[1], l)
		tempFilenames[0], tempFilenames[1] = tempFilenames[1], tempFilenames[0]
	}

	os.Rename(tempFilenames[0], outputFilename)
	os.Remove(tempFilenames[1])
}

func presort(inputFilename, outputFilename string, n, presortBatchSize int) {
	inputFile := openFile(inputFilename)
	defer inputFile.Close()

	outputFile := createFile(outputFilename)
	defer outputFile.Close()

	inputScanner := bufio.NewScanner(inputFile)

	s := make([]string, 0)
	for inputScanner.Scan() {
		s = append(s, strings.TrimSpace(inputScanner.Text()))

		if len(s) == presortBatchSize {
			dumpSortedSrings(s, outputFile)
			s = make([]string, 0)
		}
	}

	if len(s) > 0 {
		dumpSortedSrings(s, outputFile)
	}
}
