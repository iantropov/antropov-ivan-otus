package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network-4/auth"
	"social-network-4/storage"
	"social-network-4/types"
	"strings"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bearerHeader := auth.ExtractBearerAuthHeader(r.Header.Get("Authorization"))
	claims, err := auth.VerifyJWT(bearerHeader)
	if err != nil {
		fmt.Println("Failed to check JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println("CURRENT_USER_ID", claims["userId"])

	userId := strings.TrimPrefix(r.URL.Path, "/user/get/")
	userRecord, err := storage.GetUser(userId)
	if err != nil {
		fmt.Println("Failed to handle /user/get/", userId, err)
		_, ok := err.(*types.UserNotFoundError)
		if ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRecord)
}
