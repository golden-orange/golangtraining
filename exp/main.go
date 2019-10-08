package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// make connection
// see NOTE below re role of setting header on document rendering

//
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // sets page as text/html
	fmt.Fprint(w, "<h1>This is my home page</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // sets page as text/html
	fmt.Fprint(w, "<h1>Please contact me</h1>")
}

func main() {

	// 1. Most basic method, without router
	// assign page to address, send
	// http.HandleFunc("/", home)
	// http.HandleFunc("/contact", contact)

	// 2. import and initiate router
	r := mux.NewRouter()
	// assign http.HandleFunc to router variable, that will be passed as a value to ListenAndServe
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact) // URL address slice requires the "/" prefix, unless explicitly trimmed.
	// pass instead of nil, value
	http.ListenAndServe(":8021", r)

}

// NOTE: pages require explicit calls to header. The http header dictates type of file and therefore required
// protocol. Also, running the above using Println("some text to write to screen") rather than explicitly
// calling the ResponseWriter (w, "some text...") is a key final step to rendering page as html. Without the
// call to the ResponseWriter the function cannot add any content to the header. The "w" represents the
// ResponseWriter and what the ResponseWriter is or contains. Therefore,
// w = ResponseWriter = body("some text") + html document header
// The function is then sent as a value to the router - containing the client connection request and the
// document type and content to be served.

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html") // sets page as text/html
// 	fmt.Fprint(w, "<h1>This is my home page</h1>")
// }

// This function design explicitly creates page type and content information, which is inappropriate for most websites which typically
// contain larger and more pages.
// Layered abstraction is required to separate the individual information required to render the page and handle error.
// Potential abstraction clues are found where there is repetition - for example, setting Header(), initiating connection (request, response)
// and calling page content created elsewhere.
//////////////////////////////////////////////////////////////////////////////////////////////////////////
