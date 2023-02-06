package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network-3/storage"
)

func UserSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// bearerHeader := auth.ExtractBearerAuthHeader(r.Header.Get("Authorization"))
	// err := auth.VerifyJWT(bearerHeader)
	// if err != nil {
	// 	fmt.Println("Failed to check JWT token", err)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	if firstName == "" || lastName == "" {
		fmt.Println("Missed params for /user/search", firstName, lastName)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userRecords, err := storage.SearchUsers(firstName, lastName)
	if err != nil {
		fmt.Println("Failed to handle /user/search", firstName, lastName, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRecords)
}
