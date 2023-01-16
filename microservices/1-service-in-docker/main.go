package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	fmt.Println("Hello from the service in Docker!")
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)

	fmt.Println("Will serve on port", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: /health")
	w.Write([]byte(`{"status": "OK"}`))
}
