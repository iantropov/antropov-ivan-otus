package primes

import (
	"algebraic-algorithms/tester"
	"strconv"
	"testing"
)

const PRIMES_TEST_DIR = "/Users/antropov-ivan/Downloads/5.Primes/"

func TestPowerBruteForce(t *testing.T) {
	testPowerWithFiles(BruteForce, "brute-force.go", t)
}

func TestPowerOptimal(t *testing.T) {
	testPowerWithFiles(Optimal, "optimal.go", t)
}

func TestPowerEratosthenes(t *testing.T) {
	testPowerWithFiles(Eratosthenes, "eratosthenes.go", t)
}

func TestPowerEratosthenesWithBits(t *testing.T) {
	testPowerWithFiles(EratosthenesWithBits, "eratosthenes-with-bits.go", t)
}

func TestPowerLinearEratosthenes(t *testing.T) {
	testPowerWithFiles(LinearEratosthenes, "linear-eratosthenes.go", t)
}

func testPowerWithFiles(implementation Primes, testSuiteName string, t *testing.T) {
	tester.TestWithFiles(PRIMES_TEST_DIR, testSuiteName, t, func(inputs []string) []string {
		num, err := strconv.Atoi(inputs[0])
		if err != nil {
			panic(err)
		}
		outputNumber := implementation(num)
		return []string{strconv.FormatInt(int64(outputNumber), 10)}
	}, func(expected, actual string) bool {
		return expected == actual
	})
}
