package main

import (
	"github.com/bmizerany/pat"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(index))
	mux.Post("/", http.HandlerFunc(send))
	mux.Get("/confirmation", http.HandlerFunc(confirmation))

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	render(w, "./index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	//Validate Form Logic
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "./confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, datum interface{}) {
	path := filename
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err == nil {
        var contentType string

        if strings.HasSuffix(path, ".css") {
            contentType = "text/css"
        } else if strings.HasSuffix(path, ".html") {
            contentType = "text/html"
        } else if strings.HasSuffix(path, ".js") {
            contentType = "application/javascript"
        } else if strings.HasSuffix(path, ".png") {
            contentType = "image/png"
        } else if strings.HasSuffix(path, ".svg") {
            contentType = "image/svg+xml"
        } else {
            contentType = "text/css"
        }

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	}
}