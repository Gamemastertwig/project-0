// Package webface creates an HTTP interface for NPCGEN
package webface

import (
	"log"
	"net/http"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("webface/create.html", "webface/view.html"))

var validPath = regexp.MustCompile("^/(create|view)/([a-zA-Z0-9]+)$")

func loadPage(title string, body []byte) *Page {
	return &Page{Title: title, Body: body}
}

func StartHTTP(port string) {
	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("Testing", []byte("dfsdfsahjdflksjdglksjdfglksj"))
		renderTemplate(w, "view", p)
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("", []byte(""))
		renderTemplate(w, "create", p)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
