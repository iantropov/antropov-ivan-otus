package importer

import (
	"bufio"
	"fmt"
	"os"
	"social-network-2/storage"
	"social-network-2/types"
	"strconv"
	"strings"
)

func ImportPeople(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Importing users...")

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		nameParts := strings.Split(parts[0], " ")
		age, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		_, err = storage.CreateUser(types.UserRegisterParams{
			FirstName:  &nameParts[0],
			SecondName: &nameParts[1],
			Age:        &age,
			City:       parts[2],
			Password:   &parts[1],
		})

		if err != nil {
			panic(err)
		}

		counter++
		if counter%10_000 == 0 {
			fmt.Println("Added 10_000 users", counter/10_000)
		}
	}
}

// FirstName  *string `json:"first_name"`
// SecondName *string `json:"second_name"`
// Age        *int    `json:"age"`
// Biography  string  `json:"biography"`
// City       string  `json:"city"`
// Password   *string `json:"password"`
