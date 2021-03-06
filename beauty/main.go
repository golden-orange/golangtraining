package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// make function to pass as value to HandleFunc handler parameter
// use if statement in place of router. Setting up a router will require refactoring handlerFunc() into separate page functions.
// Doing so is the first level of abstraction required for simplifying and separating repetitive tasks into individual functions.

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Welcome to Beauty by Masayo!</h1>")
// 	} else if r.URL.Path == "/about" {
// 		fmt.Fprint(w, "<h1>About Beauty by Masayo!</h1>")
// 	} else if r.URL.Path == "/contact" {
// 		fmt.Fprint(w, "<h1>Contact Beauty by Masayo</h1>\n")
// 		fmt.Fprint(w, "To get in touch, please contact <a href=\"mailto:findme@scottlaing.ca\">"+
// 			"info@beautybymasayo.com</a>.")
// 	} else if r.URL.Path == "/workshops" {
// 		fmt.Fprint(w, "<h1>Workshops by Beauty by Masayo!</h1>")
// 	} else if r.URL.Path == "/headshots" {
// 		fmt.Fprint(w, "<h1>Headshots by Beauty by Masayo!</h1>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>404 Error</h1>\n")
// 		fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2> <p>Please contact us if you require further assistance.</p>")
// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	// set document type/header information
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Beauty by Masayo!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	// set document type/header information
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contact Beauty by Masayo</h1>\n")
	fmt.Fprint(w, "To get in touch, please contact <a href=\"mailto:findme@scottlaing.ca\">"+
		"info@beautybymasayo.com</a>.")
}

func about(w http.ResponseWriter, r *http.Request) {
	// set document type/header information
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>About Beauty by Masayo!</h1>")
}

func workshops(w http.ResponseWriter, r *http.Request) {
	// set document type/header information
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Workshops by Beauty by Masayo!</h1>")
}

func headshots(w http.ResponseWriter, r *http.Request) {
	// set document type/header information
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Headshots by Beauty by Masayo!</h1>")
}

// Since not using if statement, error-handling and generation of 404 page must be handled differently.
// else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>404 Error</h1>\n")
// 		fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2> <p>Please contact us if you require further assistance.</p>")
// 	}

// Since not using if statement, error-handling and generation of 404 page must be handled differently.
// else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>404 Error</h1>\n")
// 		fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2> <p>Please contact us if you require further assistance.</p>")
// 	}

func main() {

	// use gorilla mux router. using a router precludes need for if statement, as values passed to router parameters will dictate
	// which page will be rendered
	router := mux.NewRouter()

	// use HandleFunc to connection client, server, utilizing path and handler
	// http.HandleFunc("/", handlerFunc)

	// assign to router value HandleFunc + parameters for each page
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/about", about)
	router.HandleFunc("/workshops", workshops)
	router.HandleFunc("/headshots", headshots)
	// set address (port) and handler (default is nil)
	http.ListenAndServe(":8011", router)
}
