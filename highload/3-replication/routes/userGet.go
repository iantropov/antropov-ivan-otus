package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network-2/auth"
	"social-network-2/storage"
	"social-network-2/types"
	"strings"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bearerHeader := auth.ExtractBearerAuthHeader(r.Header.Get("Authorization"))
	err := auth.VerifyJWT(bearerHeader)
	if err != nil {
		fmt.Println("Failed to check JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
