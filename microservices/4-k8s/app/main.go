package main

import (
	"fmt"
	"log"
	"net/http"
	"users-app/config"
	"users-app/routes"
	"users-app/storage"
)

func HandleUserRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		routes.UserGet(w, r)
	case http.MethodPost:
		routes.UserCreate(w, r)
	case http.MethodPut:
		routes.UserUpdate(w, r)
	case http.MethodDelete:
		routes.UserDelete(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Println("Hello from the users-app!")

	storage.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/", HandleUserRoute)

	fmt.Println("Will serve on port", config.Config("PORT"))
	err := http.ListenAndServe(":"+config.Config("PORT"), mux)
	log.Fatal(err)
}
