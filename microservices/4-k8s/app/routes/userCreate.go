package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"users-app/lib"
	"users-app/storage"
	"users-app/types"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var params types.UserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !lib.AreValidUserParams(params) {
		fmt.Printf("Failed to handle POST /user : %#v\n", params)
		http.Error(w, "required fields are missed", http.StatusBadRequest)
		return
	}

	fmt.Println("Received User with params:", params)

	userId, err := storage.CreateUser(params)
	if err != nil {
		fmt.Printf("Failed to handle POST /user : %W\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userResponse := types.UserCreatedResponse{UserId: userId}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}
