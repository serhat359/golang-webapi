package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	server := http.ListenAndServe(":24672", router)
	log.Fatal(server)
}
