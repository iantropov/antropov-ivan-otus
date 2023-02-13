package routes

import (
	"fmt"
	"net/http"
	"social-network-4/auth"
	"social-network-4/storage"
	"strings"
)

func FriendDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", "PUT")
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

	userId, ok := claims["userId"].(string)
	if !ok {
		fmt.Println("Invalid userId in JWT claim", claims["userId"])
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	friendId := strings.TrimPrefix(r.URL.Path, "/friend/delete/")
	err = storage.DeleteFriend(userId, friendId)
	if err != nil {
		fmt.Println("Failed to handle /friend/delete/", friendId, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
