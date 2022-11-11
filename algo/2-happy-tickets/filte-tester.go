package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
	"testing"
)

type TestSet struct {
	hasIn, hasOut         bool
	inContent, outContent string
}

func buildTestSets(filesDir string) (map[string]*TestSet, error) {
	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		fmt.Println(err)
	}
	if len(files) == 0 {
		fmt.Printf("There are no test files in %s\n", filesDir)
		return nil, errors.New("Empty dir")
	}

	fmt.Printf("Searching test code files from %s\n", filesDir)

	testSets := make(map[string]*TestSet)

	testNameRegex := regexp.MustCompile(`\.(?:in|out)`)
	for _, f := range files {
		fileName := f.Name()
		hasInSuffix := strings.HasSuffix(fileName, ".in")
		hasOutSuffix := strings.HasSuffix(fileName, ".out")
		if !hasInSuffix && !hasOutSuffix {
			fmt.Printf("Unexpected filename %s, ignored.\n", fileName)
			continue
		}

		testName := string(testNameRegex.ReplaceAll([]byte(fileName), []byte("")))

		fileContent, readError := ioutil.ReadFile(path.Join(filesDir, fileName))
		if readError != nil {
			fmt.Printf("Failed to read file %s. Error: %v. ignored.\n", fileName, readError)
			continue
		}
		stringFileContent := strings.TrimSpace(string(fileContent))

		testSet, found := testSets[testName]
		if !found {
			testSet = &TestSet{}
			testSets[testName] = testSet
		}

		testSet.hasIn = testSet.hasIn || hasInSuffix
		testSet.hasOut = testSet.hasOut || hasOutSuffix

		if hasInSuffix {
			testSet.inContent = stringFileContent
		} else {
			testSet.outContent = stringFileContent
		}
	}

	fmt.Printf("Found %v test sets. Starting testing...\n", len(testSets))

	return testSets, nil
}

func testWithFiles(filesDir string, t *testing.T, testCallback func(string) string) {
	testSets, err := buildTestSets(filesDir)
	if err != nil {
		return
	}

	for testName, testSet := range testSets {
		t.Run(testName, func(t *testing.T) {
			actualOutput := testCallback(testSet.inContent)
			if actualOutput != testSet.outContent {
				t.Errorf("Test Failed. Expected: %s, Received: %s", testSet.outContent, actualOutput)
			}
		})
	}
}
