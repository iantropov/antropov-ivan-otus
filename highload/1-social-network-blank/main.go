package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/params"
	"social-network/storage"
)

const PORT = ":3000"

type people struct {
	Number int `json:"number"`
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

	storage.Init()

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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params params.UserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if params.FirstName == nil || params.SecondName == nil || params.Age == nil || params.Password == nil {
		http.Error(w, "required fields are missed", http.StatusBadRequest)
		return
	}

	fmt.Println("Received User with params:", params)

	userId, err := storage.CreateUser(params)
	if err != nil {
		fmt.Println("Failed to create User:", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	userResponse := UserResponse{userId}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}
