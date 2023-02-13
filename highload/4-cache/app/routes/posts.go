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

func PostCreate(w http.ResponseWriter, r *http.Request) {
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

	var params types.PostParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		fmt.Println("Invalid userId in JWT claim", claims["userId"])
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postId, err := storage.CreatePost(userId, params.Text)
	if err != nil {
		fmt.Println("Failed to handle /post/create/", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postResponse := types.PostCreateResponse{Id: postId}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postResponse)
}

func PostDelete(w http.ResponseWriter, r *http.Request) {
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

	postId := strings.TrimPrefix(r.URL.Path, "/post/delete/")
	err = storage.DeletePost(userId, postId)
	if err != nil {
		fmt.Println("Failed to handle /post/delete/", postId, err)
		_, ok := err.(*types.NotFoundError)
		if ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bearerHeader := auth.ExtractBearerAuthHeader(r.Header.Get("Authorization"))
	_, err := auth.VerifyJWT(bearerHeader)
	if err != nil {
		fmt.Println("Failed to check JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	postId := strings.TrimPrefix(r.URL.Path, "/post/get/")
	postRecord, err := storage.GetPost(postId)
	if err != nil {
		fmt.Println("Failed to handle /post/get/", postId, err)
		_, ok := err.(*types.NotFoundError)
		if ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postRecord)
}

func PostUpdate(w http.ResponseWriter, r *http.Request) {
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

	var params types.PostParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		fmt.Println("Invalid userId in JWT claim", claims["userId"])
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postId := strings.TrimPrefix(r.URL.Path, "/post/update/")
	err = storage.UpdatePost(userId, postId, params.Text)
	if err != nil {
		fmt.Println("Failed to handle /post/update/", postId, err)
		_, ok := err.(*types.NotFoundError)
		if ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
