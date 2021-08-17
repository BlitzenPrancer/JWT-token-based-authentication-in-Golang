package main

import (
	"log"
	"net/http"
)

func main() {
	// defining three handlers over here
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/referesh", Referesh)

	// starting the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
