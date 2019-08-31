package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// use glob to call all files within a directory, rather than by individual name
// reduces need to manage files individually
// utilizes many-to-few management, or lever/leverage
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
