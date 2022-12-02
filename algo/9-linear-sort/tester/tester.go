package tester

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
)

const TEST_FILENAME = "output.bin"

func GenerateFile(len int) string {
	outputFile := createFile(TEST_FILENAME)
	defer outputFile.Close()

	buf := make([]byte, 2)
	for i := 0; i < len; i++ {
		number := uint16(rand.Intn(65535))

		binary.LittleEndian.PutUint16(buf, number)
		_, error := outputFile.Write(buf)
		if error != nil {
			panic(error)
		}

		if i > 0 && i%(len/10) == 0 {
			fmt.Printf("Wrote %d/10 part.\n", i/(len/10))
		}
	}

	fmt.Println("Wrote 10/10 part.")

	outputFile.Sync()

	return outputFile.Name()
}

func ReadNumbers(filename string) []int {
	inputFile := openFile(filename)
	defer inputFile.Close()

	numbers := make([]int, 0)
	buf := make([]byte, 2)
	for {
		count, error := inputFile.Read(buf)
		if count < 2 && error != io.EOF {
			panic("Invalid read")
		} else if error == io.EOF {
			break
		}

		number := binary.LittleEndian.Uint16(buf)
		numbers = append(numbers, int(number))
	}

	return numbers
}

func createFile(filename string) *os.File {
	file, error := os.Create(filename)
	if error != nil {
		panic(error)
	}
	return file
}

func openFile(filename string) *os.File {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	return file
}
