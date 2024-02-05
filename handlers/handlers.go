package handlers

import (
	"fmt"
	"go-htmx/types"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	/*
		NOTE
		- The ParseFiles method returns a pointer to a template and an error
		- The template.Must function is a helper function that takes a pointer to a template and an error as arguments
		- If the error is not nil, the Must function panics
		- If the error is nil, the Must function returns the pointer to the template
		- The Must function is used to simplify error handling
	*/
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// tmpl := template.Must(template.ParseFiles("/index.html"))

	/*
		NOTE
			Create a map to hold the data for the template and assign it to the movies variable.
			The key is "Movies" and the value is a slice of Movies.
			The slice contains five Movie structs and each Movie struct has a Title and a Director.
			The data is hard-coded for this example and in a real-world application, the data would come from a database or another source.
	*/
	movies := map[string][]types.Movies{
		"Movies": {
			{Title: "Casablanca", Director: "Michael Curtiz"},
			{Title: "Cool Hand Luke", Director: "Stuart Rosenberg"},
			{Title: "Bullitt", Director: "Peter Yates"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The French Connection", Director: "William Friedkin"},
		},
	}

	// Execute the template and write it to the http.ResponseWriter
	if err := tmpl.Execute(w, movies); err != nil {
		// If there was an error, use the http.Error method to send a 500 internal server error response to the client
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddMovie(w http.ResponseWriter, r *http.Request) {

	// Get the form values from the request and assign them to variables
	title := r.PostFormValue("title")       // title is the name attribute of the input element in the form
	director := r.PostFormValue("director") // director is the name attribute of the input element in the form

	// Create an HTML string using the form values and a template
	htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)

	// Create a new template and parse the HTML string into it using the Parse method
	tmpl, err := template.New("t").Parse(htmlStr)
	if err != nil {
		// If there was an error, use the http.Error method to send a 500 internal server error response to the client
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and write it to the http.ResponseWriter
	tmpl.Execute(w, nil)
}
