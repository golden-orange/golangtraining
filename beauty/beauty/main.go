package main

import (
	"fmt"
	"net/http"
)

// make function to pass as value to HandleFunc handler parameter

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to Beauty by Masayo!</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>About Beauty by Masayo!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Contact Beauty by Masayo</h1>\n")
		fmt.Fprint(w, "To get in touch, please contact <a href=\"mailto:findme@scottlaing.ca\">"+
			"info@beautybymasayo.com</a>.")
	} else if r.URL.Path == "/workshops" {
		fmt.Fprint(w, "<h1>Workshops by Beauty by Masayo!</h1>")
	} else if r.URL.Path == "/headshots" {
		fmt.Fprint(w, "<h1>Headshots by Beauty by Masayo!</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 Error</h1>\n")
		fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2> <p>Please contact us if you require further assistance.</p>")
	}
}

func main() {

	// use gorilla mux router
	// use HandleFunc to connection client, server, utilizing path and handler
	http.HandleFunc("/", handlerFunc)
	// set address (port) and handler (default is nil)
	http.ListenAndServe(":8011", nil)
}
