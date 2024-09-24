package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type Page struct {
	Count int
}

func newPage() Page {
	return Page{
		Count: 1,
	}
}

// Template Rendering

var templates *template.Template

func initTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func renderTemplate(w io.Writer, name string, data interface{}) error {
	return templates.ExecuteTemplate(w, name, data)
}

// Handlers

func indexHandler(page *Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := renderTemplate(w, "index", page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	initTemplates()
	page := newPage()

	http.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("GET /", indexHandler(&page))
	http.HandleFunc("GET /counter/increase", func(w http.ResponseWriter, r *http.Request) {
		page.Count++
		renderTemplate(w, "count", page)
	})
	http.HandleFunc("GET /counter/decrease", func(w http.ResponseWriter, r *http.Request) {
		page.Count--
		renderTemplate(w, "count", page)
	})

	log.Println("Server started at :1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}
