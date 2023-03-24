package main

import (
	"fmt"
	"log"
	"net/http"
	"social-network-5/config"
	"social-network-5/routes"
)

func main() {
	fmt.Println("Hello from the social network 4!")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", routes.Login)

	mux.HandleFunc("/user/register", routes.UserRegister)
	mux.HandleFunc("/user/get/", routes.UserGet)
	mux.HandleFunc("/user/search/", routes.UserSearch)

	mux.HandleFunc("/whoami", routes.Whoami)

	mux.HandleFunc("/friend/set/", routes.FriendSet)
	mux.HandleFunc("/friend/delete/", routes.FriendDelete)

	mux.HandleFunc("/post/create", routes.PostCreate)
	mux.HandleFunc("/post/delete/", routes.PostDelete)
	mux.HandleFunc("/post/update/", routes.PostUpdate)
	mux.HandleFunc("/post/get/", routes.PostGet)
	mux.HandleFunc("/post/feed", routes.PostFeed)

	mux.HandleFunc("/dialog/", routes.Dialog)

	fmt.Println("Will serve on port", config.Config("PORT"))
	err := http.ListenAndServe(":"+config.Config("PORT"), mux)
	log.Fatal(err)
}
