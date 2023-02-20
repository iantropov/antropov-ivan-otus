package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"users-app/lib"
	"users-app/storage"
	"users-app/types"
)

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	userIdStr := strings.TrimPrefix(r.URL.Path, "/user/")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		fmt.Printf("Failed to handle PUT /user/%s: %W\n", userIdStr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var params types.UserParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !lib.AreValidUserParams(params) {
		fmt.Printf("Failed to handle PUT /user : %W\n", err)
		http.Error(w, "required fields are missed", http.StatusBadRequest)
		return
	}

	fmt.Println("Received User with params:", params)

	err = storage.UpdateUser(userId, params)
	if err != nil {
		fmt.Printf("Failed to handle PUT /user : %W\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
