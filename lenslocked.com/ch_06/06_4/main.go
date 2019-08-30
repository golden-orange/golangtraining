package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/ch_06/06_4/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8011", r)
}

// bootstrap builds the same layout for every view to which it is attached.
