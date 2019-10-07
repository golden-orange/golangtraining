// next step in implementing routers: using with gorilla mux
// replace if statements with calls to unique page handlers via the mux server
// this eliminates the handlerFunc function from if statement method
package main

import (
	"github.com/gorilla/mux"
	// import gorilla mux router

	"fmt"
	"net/http"
)

// make function to pass as value to HandleFunc handler parameter

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Beauty by Masayo!</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contact Beauty by Masayo</h1>\n")
	fmt.Fprint(w, "To get in touch, please contact <a href=\"mailto:findme@scottlaing.ca\">"+
		"info@beautybymasayo.com</a>.")
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>About Beauty by Masayo!</h1>")
}

func workshops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Workshops by Beauty by Masayo!</h1>")
}

func headshots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Headshots by Beauty by Masayo!</h1>")
}

func statusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Error</h1>\n")
	fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2> <p>Please contact us if you require further assistance.</p>")
}

func main() {
	var handler http.Handler
	if handler == nil {
		handler = http.NotFoundHandler()
	}

	// use gorilla mux router
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/about", about)
	r.HandleFunc("/workshops", workshops)
	r.HandleFunc("/headshots", headshots)

	var n http.Handler = http.HandlerFunc(statusNotFound)
	r.NotFoundHandler = n

	// set address (port) and handler (default is nil)
	http.ListenAndServe(":8011", r)
}
