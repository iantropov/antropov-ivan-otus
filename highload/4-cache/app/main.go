package main

import (
	"fmt"
	"log"
	"net/http"
	"social-network-4/config"
	"social-network-4/routes"
	"social-network-4/storage"
)

func main() {
	fmt.Println("Hello from the social network 4!")

	storage.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/user/register", routes.UserRegister)
	mux.HandleFunc("/login", routes.Login)
	mux.HandleFunc("/user/get/", routes.UserGet)
	mux.HandleFunc("/user/search/", routes.UserSearch)

	fmt.Println("Will serve on port", config.Config("PORT"))
	err := http.ListenAndServe(":"+config.Config("PORT"), mux)
	log.Fatal(err)
}
