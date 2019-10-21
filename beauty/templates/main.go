// go/src/beauty/templates/

// OVERVIEW:
// Review course CH 6 - Views

package main

import (
	"fmt"
	"html/template"

	"github.com/gorilla/mux"

	"net/http"
)

// first abstraction layer - place template body in views documents, then parse to template inside main
// import template global variable (this will be replaced later)
var homeTemplate *template.Template
var contactTemplate *template.Template

func statusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Error</h1>\n")
	fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2> \n <p>Please contact us if you require further assistance.</p>")
}

// requires a page that handles connection, content-call to template and will be sent as a value to HandlerFunc parameter
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	var err error

	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}

	// use gorilla mux router
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	var n http.Handler = http.HandlerFunc(statusNotFound)
	r.NotFoundHandler = n

	// set address (port) and handler (default is nil)
	http.ListenAndServe(":8011", r)
}
