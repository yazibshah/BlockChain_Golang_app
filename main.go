package main

import "github.com/gin-gonic/gin"

import (
	"log"
	"net/http"
)

func main() {
	router := createRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Define your routes here
	router.HandleFunc("/", handleRoot)

	return router
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Implement the root handler logic here
	w.Write([]byte("Hello, World!"))
}
