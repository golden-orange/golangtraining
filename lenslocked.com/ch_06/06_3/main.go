package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Get in touch!</h1>")
	err := contactTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layout/footer.gohtml",
	)

	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layout/footer.gohtml",
	)

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8010", r)
}
