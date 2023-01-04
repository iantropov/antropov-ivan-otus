package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":3000"

func main() {
	fmt.Println("Hello from the social network!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/bye", handleBye)

	fmt.Println("Will serve on port", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: Hello")
	w.Write([]byte("Hello"))
}

func handleBye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: Bye")
	w.Write([]byte("Bye"))
}
