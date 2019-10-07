// update router method, using violetear

package main

import (
	"fmt"
	"net/http"

	// import violetear router
	"github.com/nbari/violetear"
)

func main() {

	// create handler to be passed as value to HandleFunc
	// create individual variables, each assigned to an address
	home := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Welcome to tiny houses news site!</h1>")
	}

	contact := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Contact Tiny Houses</h1>\n")
		fmt.Fprint(w, "To get in touch, please contact <a href=\"mailto:findme@scottlaing.ca\">"+
			"contact@tinyhouses.com</a>.")
	}
	// notFound := func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "<h1>404 Error</h1>\n")
	// 	fmt.Fprint(w, "<h2>We could not find the page you were seeking.</h2>"+"<p>Please contact us if you require further assistance.</p>")
	// }
	fitOverFifty := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Welcome to Fit Over Fifty!</h1>"+
			"<p>For those seeking exceptional quality after the big Five - Oh</p>")
	}
	jewishQuestion := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Exploring questions governments and big corps prefer to ignore</h1>")
	}
	crimeCity := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>A non-mainstream look at actual crime data and meaning</h1>")
	}
	// set new instance of violetear router
	vr := violetear.New()
	vr.HandleFunc("/", home)
	vr.HandleFunc("/fitoverfifty/", fitOverFifty)
	vr.HandleFunc("/jq/", jewishQuestion)
	vr.HandleFunc("/crimecity/", crimeCity)
	vr.HandleFunc("/contact", contact)

	// set address (port) and use default value for handler
	http.ListenAndServe(":8013", vr)
}
