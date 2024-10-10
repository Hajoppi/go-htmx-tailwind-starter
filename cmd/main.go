package main

import (
	"htmx-starter/internal"
	"log"
	"net/http"
)

func main() {
	app := internal.NewApp()
	log.Println("Server started at :1337")
	log.Fatal(http.ListenAndServe(":1337", app))
}
