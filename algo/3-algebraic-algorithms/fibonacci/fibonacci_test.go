package fibonacci

import (
	"algebraic-algorithms/tester"
	"strconv"
	"testing"
)

const FIBONACCI_TEST_DIR = "/Users/antropov-ivan/Downloads/4.Fibo/"

func TestPowerRecursive(t *testing.T) {
	testPowerWithFiles(Recursive, "recursive.go", t)
}

func TestPowerIterative(t *testing.T) {
	testPowerWithFiles(Iterative, "iterative.go", t)
}

func TestPowerGoldenRatio(t *testing.T) {
	testPowerWithFiles(GoldenRatio, "golden-ratio.go", t)
}

func TestPowerMatrix(t *testing.T) {
	testPowerWithFiles(Matrix, "matrix.go", t)
}

func testPowerWithFiles(implementation Fibonacci, testSuiteName string, t *testing.T) {
	tester.TestWithFiles(FIBONACCI_TEST_DIR, testSuiteName, t, func(inputs []string) []string {
		num, err := strconv.Atoi(inputs[0])
		if err != nil {
			panic(err)
		}
		outputNumber := implementation(num)
		return []string{outputNumber.String()}
	}, func(expected, actual string) bool {
		return expected == actual
	})
}
