package power

import (
	"algebraic-algorithms/tester"
	"math"
	"strconv"
	"testing"
)

const POWER_TEST_DIR = "/Users/antropov-ivan/Downloads/3.Power/"

func TestPowerIterative(t *testing.T) {
	testPowerWithFiles(Iterative, "iterative.go", t)
}

func TestPowerSuboptimal(t *testing.T) {
	testPowerWithFiles(Suboptimal, "suboptimal.go", t)
}

func TestPowerOptimal(t *testing.T) {
	testPowerWithFiles(Optimal, "optimal.go", t)
}

func testPowerWithFiles(implementation Power, testSuiteName string, t *testing.T) {
	tester.TestWithFiles(POWER_TEST_DIR, testSuiteName, t, func(inputs []string) []string {
		num, err := strconv.ParseFloat(inputs[0], 64)
		if err != nil {
			panic(err)
		}
		pow, err := strconv.Atoi(inputs[1])
		if err != nil {
			panic(err)
		}
		outputNumber := implementation(num, pow)
		return []string{strconv.FormatFloat(outputNumber, 'f', 12, 64)}
	}, func(expected, actual string) bool {
		expectedFloat, err := strconv.ParseFloat(expected, 64)
		if err != nil {
			panic(err)
		}
		actualFloat, err := strconv.ParseFloat(actual, 64)
		if err != nil {
			panic(err)
		}
		return math.Abs(expectedFloat-actualFloat) < 0.0000001
	})
}
