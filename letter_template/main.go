package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template using a constant
	const letter = `
		Dear {{.Name}},
		{{if .AcceptedInvitation}}
		I am glad you are able to attend the annual Gathering in the Woods.	
		{{with .Trait -}}
		Please be sure to leave all {{.}} at home as it won't be needed, nor tolerated.
		{{- end}}
		{{- else}}
		I am disappointed you are unable to attend the annual Gathering.
		Please be sure we will chant and plot in your absence.
		{{end}}
		Your Friend,
		Jose, the cult leader
		`

	// data to be entered into template
	type potentialRecruit struct {
		Name, Trait        string
		AcceptedInvitation bool
	}

	// list of potential recruits and current status
	var potentialRecruits = []potentialRecruit{
		{"Bob", "guile", true},
		{"Betty", "intelligence", true},
		{"Tom", "connection to reality", false},
		{"Mick", "whatever", false},
	}

	// create a new template and parse the letter into it
	l := template.Must(template.New("letter").Parse(letter))

	// execute the letter template for each potential recruit
	for _, r := range potentialRecruits {
		err := l.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
