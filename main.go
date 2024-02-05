package main

import (
	"go-htmx/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Home)               // Home is the handler for the root route
	http.HandleFunc("/add-movie/", handlers.AddMovie) // AddMovie is the handler for the /add-movie/ route

	log.Println("Server started on: 8000")

	// we use the http.ListenAndServe method to start the server and listen for incoming requests on port 8000,
	// we also pass nil as the second argument to use the default router for the server.
	// If there is an error starting the server, we log the error to the console.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
