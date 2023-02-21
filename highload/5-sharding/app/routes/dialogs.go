package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network-5/auth"
	"social-network-5/storage"
	"social-network-5/types"
	"strings"
)

func Dialog(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/dialog/")
	if strings.HasSuffix(path, "/send") {
		dialogSendMessage(w, r)
	} else if strings.HasSuffix(path, "/list") {
		dialogListMessages(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func dialogSendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
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

	var params types.MessageParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipientId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/dialog/"), "/send")
	dialogId, err := storage.GetDialogId(userId, recipientId)
	if err != nil {
		fmt.Println("Failed to get dialog", err)
		return
	}
	if dialogId == "" {
		dialogId, err = storage.CreateDialog(fmt.Sprintf("A cozy chat for %s and %s", userId, recipientId))
		if err != nil {
			fmt.Println("Failed to create dialog", err)
			return
		}
	}

	err = storage.CreateMessage(userId, recipientId, dialogId, params.Text)
	if err != nil {
		fmt.Println("Failed to create message", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func dialogListMessages(w http.ResponseWriter, r *http.Request) {
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

	userId, ok := claims["userId"].(string)
	if !ok {
		fmt.Println("Invalid userId in JWT claim", claims["userId"])
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	recipientId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/dialog/"), "/list")
	dialogId, err := storage.GetDialogId(userId, recipientId)
	if err != nil {
		fmt.Println("Failed to get dialog", err)
		return
	}
	if dialogId == "" {
		json.NewEncoder(w).Encode([]types.MessageRecord{})
		return
	}

	messages, err := storage.GetMessages(dialogId)
	if err != nil {
		fmt.Println("Failed to get messages", err)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
