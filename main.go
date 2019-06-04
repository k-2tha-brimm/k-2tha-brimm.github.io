package main

import (
	"net/http"
)

func main() {

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir("./")))

    http.ListenAndServe(":3000", nil)
}