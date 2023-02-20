package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"users-app/storage"
	"users-app/types"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	userIdStr := strings.TrimPrefix(r.URL.Path, "/user/")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		fmt.Printf("Failed to handle GET /user/%s: %W\n", userIdStr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userRecord, err := storage.GetUser(userId)
	if err != nil {
		fmt.Printf("Failed to handle GET /user/%d: %W\n", userId, err)
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
