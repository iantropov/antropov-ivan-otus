package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":3000"

type people struct {
	Number int `json:"number"`
}

type UserParams struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int    `json:"age"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password"`
}

type UserResponse struct {
	UserId string `json:"user_id"`
}

func main() {
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)

	fmt.Println("Hello from the social network!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/bye", handleBye)
	mux.HandleFunc("/user/register", handleUserRegister)

	fmt.Println("Will serve on port", PORT)
	err = http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: Hello")
	w.Write([]byte("Hello"))
}

func handleBye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: Bye")
	w.Write([]byte("Bye"))
}

func handleUserRegister(w http.ResponseWriter, r *http.Request) {
	var params UserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Received User with params:", params)

	userResponse := UserResponse{"e4d2e6b0-cde2-42c5-aac3-0b8316f21e58"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}
