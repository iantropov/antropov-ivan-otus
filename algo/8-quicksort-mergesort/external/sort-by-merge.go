package external

import (
	"bufio"
	"fmt"
	"os"
)

const FILE_IS_OVER = -1

type Reader struct {
	scanner *bufio.Scanner
	pos     int
}

func SortByMerge(inputFilename, outputFilename string, n int) {
	tempFilenames := []string{"temp-0.text", "temp-1.text"}
	singleMergePass("input.txt", tempFilenames[0], 1)

	for l := 2; l < n; l *= 2 {
		singleMergePass(tempFilenames[0], tempFilenames[1], l)
		tempFilenames[0], tempFilenames[1] = tempFilenames[1], tempFilenames[0]
	}

	os.Rename(tempFilenames[0], outputFilename)
	os.Remove(tempFilenames[1])
}

func singleMergePass(inputFilename, outputFilename string, l int) {
	outputFile, error := os.Create(outputFilename)
	if error != nil {
		panic(error)
	}
	defer outputFile.Close()

	files := []*os.File{openFile(inputFilename), openFile(inputFilename)}
	defer files[0].Close()
	defer files[1].Close()

	readers := []*Reader{buildReader(files[0]), buildReader(files[1])}
	rewindReader(readers[1], l)

	values := []int{readNextNumberUntil(readers[0], l), readNextNumberUntil(readers[1], l)}

	for isThereAvailableValue(values) {
		for i := 0; i < 2*l && isThereAvailableValue(values); i++ {
			minIndex := findMinIndex(values)
			_, error = outputFile.WriteString(fmt.Sprint(values[minIndex]) + "\n")
			if error != nil {
				panic(error)
			}
			values[minIndex] = readNextNumberUntil(readers[minIndex], l)
		}

		rewindReader(readers[0], l)
		rewindReader(readers[1], l)
		values[0] = readNextNumberUntil(readers[0], l)
		values[1] = readNextNumberUntil(readers[1], l)
	}
}

func buildReader(file *os.File) *Reader {
	reader := new(Reader)
	reader.scanner = bufio.NewScanner(file)
	return reader
}

func openFile(filename string) *os.File {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	return file
}

func readNextNumberUntil(reader *Reader, l int) int {
	if reader.pos < l && reader.scanner.Scan() {
		reader.pos++
		return readNumber(reader.scanner)
	} else {
		return FILE_IS_OVER
	}
}

func rewindReader(reader *Reader, l int) {
	for i := 0; i < l; i++ {
		reader.scanner.Scan()
	}
	reader.pos = 0
}
