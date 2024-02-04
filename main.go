package main

import (
	"fmt"
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

	h2 := func(w http.ResponseWriter, r *http.Request) {

		// Get the form values from the request body
		// The form values are sent as part of the request body
		// The request body is read by the server and the form values are extracted
		// The form values are then made available to the server as part of the request
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)

		// HX-Request header is set to "true" when the request is made by HTMX
		// log.Println(r.Header.Get("HX-Request"))
	}

	// Create a handler which uses the function
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Println("Server started on: http://localhost:8000")

	// Start the server on port 8000
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
