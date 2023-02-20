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
		_, ok := err.(*types.NotFoundError)
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

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params types.UserRegisterParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if params.FirstName == nil || params.SecondName == nil || params.Age == nil || params.Password == nil {
		http.Error(w, "required fields are missed", http.StatusBadRequest)
		return
	}

	fmt.Println("Received User with params:", params)

	userId, err := storage.CreateUser(params)
	if err != nil {
		fmt.Println("Failed to handle /user/register:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userResponse := types.UserRegisterResponse{UserId: userId}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func UserSearch(w http.ResponseWriter, r *http.Request) {
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
