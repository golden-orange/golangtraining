package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/ch_06/06_7/views"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

// adding a helper function that will panic at any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8012", r)
}

//lenslocked.com ch_06 06_7
