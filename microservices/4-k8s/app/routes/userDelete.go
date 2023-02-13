package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"users-app/storage"
	"users-app/types"
)

func UserDelete(w http.ResponseWriter, r *http.Request) {
	userIdStr := strings.TrimPrefix(r.URL.Path, "/user/")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		fmt.Printf("Failed to handle DELETE /user/%s: %W\n", userIdStr, err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	err = storage.DeleteUser(userId)
	if err != nil {
		fmt.Printf("Failed to handle DELETE /user/%d: %W\n", userId, err)
		_, ok := err.(*types.UserNotFoundError)
		if ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
