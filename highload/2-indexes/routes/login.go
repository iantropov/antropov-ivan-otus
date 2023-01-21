package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network-2/auth"
	"social-network-2/storage"
	"social-network-2/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	token, err := auth.GenerateJWT(userRecord.Id)
	if err != nil {
		fmt.Println("Failed to handle generte JWT:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := types.UserLoginResponse{Token: token}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
