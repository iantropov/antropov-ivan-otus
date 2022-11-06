package main

import (
	"strconv"
	"testing"
)

const TEST_DIR = "/Users/antropov-ivan/Downloads/A01_Счастливые_билеты-1801-057a77/1.Tickets"

func Test_happyTicketsWithSeniorSolution(t *testing.T) {
	testWithFiles(TEST_DIR, t, func(input string) string {
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		outputNumber := FindHappyTicketsAsSenior(inputNumber)
		return strconv.Itoa(outputNumber)
	})
}
