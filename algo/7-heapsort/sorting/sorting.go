package sorting

import (
	"heapsort/tester"
	"strconv"
	"strings"
	"testing"
)

const RANDOM_TESTS_DIR = "/Users/antropov-ivan/Downloads/sorting-tests/0.random/"
const DIGITS_TESTS_DIR = "/Users/antropov-ivan/Downloads/sorting-tests/1.digits/"
const SORTED_TESTS_DIR = "/Users/antropov-ivan/Downloads/sorting-tests/2.sorted/"
const REVERS_TESTS_DIR = "/Users/antropov-ivan/Downloads/sorting-tests/3.revers/"

type Sort func([]int) []int

func Swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func TestSorting(sort Sort, t *testing.T) {
	array := []int{1, 23, 4, 7, 2, 9, 0, 2, 89, 41}
	expectedSortedArray := []int{0, 1, 2, 2, 4, 7, 9, 23, 41, 89}

	actualSortedArray := sort(array)

	for i := 0; i < len(array); i++ {
		if actualSortedArray[i] != expectedSortedArray[i] {
			t.Errorf("Incorrect actual value - %d. Expected - %d\n", actualSortedArray[i], expectedSortedArray[i])
		}
	}
}

func TestSortingWithRandomFiles(sort Sort, t *testing.T) {
	testSortingWithFiles(sort, RANDOM_TESTS_DIR, "random", t)
}

func TestSortingWithDigitsFiles(sort Sort, t *testing.T) {
	testSortingWithFiles(sort, DIGITS_TESTS_DIR, "digits", t)
}

func TestSortingWithSortedFiles(sort Sort, t *testing.T) {
	testSortingWithFiles(sort, SORTED_TESTS_DIR, "sorted", t)
}

func TestSortingWithReversFiles(sort Sort, t *testing.T) {
	testSortingWithFiles(sort, REVERS_TESTS_DIR, "revers", t)
}

func testSortingWithFiles(sort Sort, testsDir, testSuiteName string, t *testing.T) {
	tester.TestWithFiles(testsDir, testSuiteName, t, func(inputs []string) []string {
		num, err := strconv.Atoi(inputs[0])
		if err != nil {
			panic(err)
		}
		a := make([]int, num)
		numStrings := strings.Split(inputs[1], " ")
		if len(numStrings) != num {
			panic("invalid count of numbers")
		}
		for i, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic("invalid number string: " + numString)
			}
			a[i] = num
		}

		sorted := sort(a)

		output := make([]string, len(sorted))
		for i, num := range sorted {
			output[i] = strconv.FormatInt(int64(num), 10)
		}

		return []string{strings.Join(output, " ")}
	}, func(expected, actual string) bool {
		return expected == actual
	})
}
