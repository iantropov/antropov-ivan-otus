package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/storage"
	"social-network/types"
	"strings"
)

const PORT = ":3000"

type people struct {
	Number int `json:"number"`
}

type UserResponse struct {
	UserId string `json:"user_id"`
}

func main() {
	fmt.Println("Hello from the social network!")

	storage.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", handleUserRegister)
	mux.HandleFunc("/login", handleUserLogin)
	mux.HandleFunc("/user/get/", handleUserGet)

	fmt.Println("Will serve on port", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}

func handleUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params types.UserRegisterParams
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
		fmt.Println("Failed to handle /user/register:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userResponse := UserResponse{userId}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func handleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params types.UserLoginParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if params.Id == nil || params.Password == nil {
		http.Error(w, "required fields are missed", http.StatusBadRequest)
		return
	}

	userRecord, err := storage.LoginUser(*params.Id, *params.Password)
	if err != nil {
		fmt.Println("Failed to handle /login:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRecord)
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := strings.TrimPrefix(r.URL.Path, "/user/get/")
	userRecord, err := storage.GetUser(userId)
	if err != nil {
		fmt.Println("Failed to handle /get/user/", userId, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRecord)
}
