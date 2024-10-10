package internal

import "net/http"

type App struct {
	page *Page
	mux  *http.ServeMux
}

func NewApp() *App {
	page := newPage()
	initTemplates()

	mux := http.NewServeMux()
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.HandleFunc("GET /counter/increase", page.increaseHandler)
	mux.HandleFunc("GET /counter/decrease", page.decreaseHandler)
	mux.HandleFunc("GET /", page.indexHandler)
	return &App{
		page: page,
		mux:  mux,
	}
}

// ServeHTTP handles HTTP requests for both Lambda and native HTTP servers.
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.mux.ServeHTTP(w, r)
}
