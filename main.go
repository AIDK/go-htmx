package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Movies struct {
	Title    string
	Director string
}

func main() {

	// Create a function which handles the request and response
	h1 := func(w http.ResponseWriter, r *http.Request) {

		// Create a template
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Create a map (dictionary) to hold the data for the template (the films data)
		// The key is "Movies" and the value is a slice of Movies
		// The slice contains three Movies
		// NOTE: The data can be fetched from a database or any other source
		films := map[string][]Movies{
			"Movies": {
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
		/*
		   NOTE
		   The form values are sent as part of the request body
		   The request body is read by the server and the form values are extracted
		   The form values are then made available to the server as part of the request
		*/
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		// we can use the form values to create an HTML string
		// and send it back to the client
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)

		// Create a template and execute it with the HTML string
		tmpl, err := template.New("t").Parse(htmlStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Execute the template and pass the HTML string to it
		tmpl.Execute(w, nil)

		// this header is set to "true" when the request is made by HTMX
		// which can be used to indicate whether the request is made by HTMX or not
		// log.Println(r.Header.Get("HX-Request"))
	}

	// Create a handler which uses the function
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-movie/", h2)

	log.Println("Server started on: 8000")

	// Start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
