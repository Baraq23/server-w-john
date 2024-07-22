package handlers

import (
	// "fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

type Info struct {
	Content string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// data2 := Info {
	// 	content: "Different content",
	// }

	tmpl, _ := template.ParseFiles(filepath.Join("templates", "home.html"))
	data1 := Info{
		Content: "this is the content",
	}
	tmpl.Execute(w, data1)
}
