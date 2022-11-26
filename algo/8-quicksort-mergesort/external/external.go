package external

import (
	"math/rand"
	"os"
	"strconv"
)

func GenerateTextFile(name string, n, t int) error {
	file, error := os.Create(name)
	if error != nil {
		return error
	}

	defer file.Close()

	var randomNumber uint32
	for i := 0; i < n; i++ {
		randomNumber = uint32(rand.Intn(t))
		_, error := file.WriteString(strconv.FormatInt(int64(randomNumber), 10) + "\n")
		if error != nil {
			return error
		}
	}

	error = file.Sync()
	if error != nil {
		return error
	}

	return nil
}
