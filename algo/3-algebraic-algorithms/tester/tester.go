package tester

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
	fileName              string
	inContent, outContent []string
}

func buildTestSets(filesDir, testSuteName string) (map[string]*TestSet, error) {
	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		fmt.Println(err)
	}
	if len(files) == 0 {
		fmt.Printf("[#%s] There are no test files in %s\n", testSuteName, filesDir)
		return nil, errors.New("Empty dir")
	}

	fmt.Printf("[#%s] Searching test code files from %s\n", testSuteName, filesDir)

	testSets := make(map[string]*TestSet)

	testNameRegex := regexp.MustCompile(`\.(?:in|out)`)
	for _, f := range files {
		fileName := f.Name()
		hasInSuffix := strings.HasSuffix(fileName, ".in")
		hasOutSuffix := strings.HasSuffix(fileName, ".out")
		if !hasInSuffix && !hasOutSuffix {
			fmt.Printf("[#%s] Unexpected filename %s, ignored.\n", testSuteName, fileName)
			continue
		}

		testName := string(testNameRegex.ReplaceAll([]byte(fileName), []byte("")))

		fileContent, readError := ioutil.ReadFile(path.Join(filesDir, fileName))
		if readError != nil {
			fmt.Printf("[#%s] Failed to read file %s. Error: %v. ignored.\n", testSuteName, fileName, readError)
			continue
		}

		fileStrings := strings.Split(string(fileContent), "\n")
		for i, fileString := range fileStrings {
			fileStrings[i] = strings.TrimSpace(fileString)
		}

		testSet, found := testSets[testName]
		if !found {
			testSet = &TestSet{}
			testSets[testName] = testSet
		}

		testSet.fileName = fileName
		testSet.hasIn = testSet.hasIn || hasInSuffix
		testSet.hasOut = testSet.hasOut || hasOutSuffix

		if hasInSuffix {
			testSet.inContent = fileStrings
		} else {
			testSet.outContent = fileStrings
		}
	}

	fmt.Printf("[#%s] Found %v test sets. Starting testing...\n", testSuteName, len(testSets))

	return testSets, nil
}

func TestWithFiles(filesDir string, testSuteName string, t *testing.T, testCallback func([]string) []string, compareCallback func(string, string) bool) {
	testSets, err := buildTestSets(filesDir, testSuteName)
	if err != nil {
		return
	}

	for testName, testSet := range testSets {
		t.Run(testName, func(t *testing.T) {
			actualOutput := testCallback(testSet.inContent)
			for i := range actualOutput {
				if !compareCallback(testSet.outContent[i], actualOutput[i]) {
					t.Errorf("[#%s] [.%s] Test Failed. Expected: %s, Received: %s", testSuteName, testSet.fileName, testSet.outContent, actualOutput)
				}
			}
		})
	}
}
