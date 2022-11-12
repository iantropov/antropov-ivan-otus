package main

import (
	"algebraic-algorithms/power"
	"strconv"
	"testing"
)

const POWER_TEST_DIR = "/Users/antropov-ivan/Downloads/3.Power/"

func TestPower(t *testing.T) {
	testPowersWithFiles([]power.Power{
		power.Iterative,
		// power.Suboptimal,
		// power.Optimal,
	}, t)
}

func testPowersWithFiles(implementations []power.Power, t *testing.T) {
	for _, implentation := range implementations {
		testPowerWithFiles(implentation, t)
	}
}

func testPowerWithFiles(implementation power.Power, t *testing.T) {
	testWithFiles(POWER_TEST_DIR, t, func(inputs []string) []string {
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
	})
}
