package main

import (
	"fmt"
	"social-network-2/importer"
	"social-network-2/storage"
)

func main() {
	fmt.Println("Hello from the social network!")

	storage.Init()

	importer.ImportPeople("./people.csv")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/user/register", routes.UserRegister)
	// mux.HandleFunc("/login", routes.Login)
	// mux.HandleFunc("/user/get/", routes.UserGet)
	// mux.HandleFunc("/user/search/", routes.UserSearch)

	// fmt.Println("Will serve on addr", config.Config("ADDR"))
	// err := http.ListenAndServe(config.Config("ADDR"), mux)
	// log.Fatal(err)
}
