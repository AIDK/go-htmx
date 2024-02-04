package main

import (
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	// Create a function which handles the request and response
	h1 := func(w http.ResponseWriter, r *http.Request) {
		// Create a template
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Director: "Michael Curtiz"},
				{Title: "Cool Hand Luke", Director: "Stuart Rosenberg"},
				{Title: "Bullitt", Director: "Peter Yates"},
			},
		}

		// Execute the template (pass the films data to the template)
		if err := tmpl.Execute(w, films); err != nil {
			// If there is an error, print it to the console
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// Create a handler which uses the function
	http.HandleFunc("/", h1)

	log.Println("Server started on: http://localhost:8000")

	// Start the server on port 8000
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
