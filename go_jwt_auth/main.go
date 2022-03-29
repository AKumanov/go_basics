package main

import (
	"log"
	"net/http"
)

func main() {
	// declare main functions
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	// start server
	log.Fatal(http.ListenAndServe(":8080", nil))

}
