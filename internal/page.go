package internal

import "net/http"

type Page struct {
	Count int
}

func newPage() *Page {
	return &Page{
		Count: 1,
	}
}

func (page *Page) indexHandler(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "index", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (page *Page) increaseHandler(w http.ResponseWriter, r *http.Request) {
	page.Count++
	renderTemplate(w, "count", page)
}

func (page *Page) decreaseHandler(w http.ResponseWriter, r *http.Request) {
	page.Count--
	renderTemplate(w, "count", page)
}
