package external

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"quicksort-mergesort/quick"
	"strconv"
	"strings"
)

const TEMP_FILE_IS_OVER = -1

func SortBySplit(inputFilename, outputFilename string, n, t int) {
	tempFilenames := generateTempFiles(inputFilename, n, t)
	tLen := len(tempFilenames)
	tempFiles := make([]*os.File, tLen)
	tempScanners := make([]*bufio.Scanner, tLen)
	tempValues := make([]int, tLen)
	outputFile, error := os.Create(outputFilename)
	if error != nil {
		panic(error)
	}

	defer outputFile.Close()

	for i := 0; i < tLen; i++ {
		tempFiles[i], error = os.Open(tempFilenames[i])
		if error != nil {
			panic(error)
		}

		tempScanners[i] = bufio.NewScanner(tempFiles[i])
		tempValues[i] = TEMP_FILE_IS_OVER
		if tempScanners[i].Scan() {
			tempValues[i] = readNumber(tempScanners[i])
		}
	}

	for isThereAvailableValue(tempValues) {
		minIndex := findMinIndex(tempValues)
		_, error = outputFile.WriteString(fmt.Sprint(tempValues[minIndex]) + "\n")
		if error != nil {
			panic(error)
		}

		if tempScanners[minIndex].Scan() {
			tempValues[minIndex] = readNumber(tempScanners[minIndex])
		} else {
			tempValues[minIndex] = TEMP_FILE_IS_OVER
		}
	}

	for i := 0; i < tLen; i++ {
		tempFiles[i].Close()
		os.Remove(tempFilenames[i])
	}
}

func generateTempFiles(inputFilename string, n, t int) []string {
	numbersPerFile := int(math.Ceil(float64(n) / float64(t)))

	inputFile, error := os.Open(inputFilename)
	if error != nil {
		panic(error)
	}

	inputScanner := bufio.NewScanner(inputFile)

	defer inputFile.Close()

	linesCounter := 0
	filesCounter := 0
	s := make([]string, 0)
	tempFilenames := make([]string, 0)

	for inputScanner.Scan() {
		s = append(s, strings.TrimSpace(inputScanner.Text()))
		linesCounter++

		if linesCounter == numbersPerFile {
			tempFilenames = append(tempFilenames, dumpStrings(s, filesCounter))
			filesCounter++

			s = make([]string, 0)
			linesCounter = 0
		}
	}

	if linesCounter > 0 {
		tempFilenames = append(tempFilenames, dumpStrings(s, filesCounter))
	}

	return tempFilenames
}

func dumpStrings(s []string, filesCounter int) string {
	tempFilename := "temp-" + fmt.Sprint(filesCounter) + ".txt"
	tempFile, error := os.Create(tempFilename)
	if error != nil {
		panic(error)
	}

	sorted := sortStrings(s)
	tempFile.WriteString(strings.Join(sorted, "\n"))
	tempFile.Sync()
	tempFile.Close()

	return tempFilename
}

func sortStrings(a []string) []string {
	numbers := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		number, error := strconv.Atoi(a[i])
		if error != nil {
			panic(error)
		}
		numbers[i] = number
	}
	sortedNumbes := quick.Sort(numbers)
	sortedStrings := make([]string, len(a))
	for i, num := range sortedNumbes {
		sortedStrings[i] = strconv.FormatInt(int64(num), 10)
	}

	return sortedStrings
}

func readNumber(scanner *bufio.Scanner) int {
	value, error := strconv.Atoi(scanner.Text())
	if error != nil {
		panic(error)
	}
	return value
}

func isThereAvailableValue(values []int) bool {
	for _, value := range values {
		if value != TEMP_FILE_IS_OVER {
			return true
		}
	}
	return false
}

func findMinIndex(values []int) int {
	minIndex := -1
	minValue := math.MaxInt64
	for i, value := range values {
		if value != -1 && value < minValue {
			minIndex = i
			minValue = value
		}
	}
	return minIndex
}
