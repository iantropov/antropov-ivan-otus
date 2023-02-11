package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network-3/storage"
	"social-network-3/types"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
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

	userResponse := types.UserRegisterResponse{UserId: userId}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}
