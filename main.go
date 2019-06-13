package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"net/smtp"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", send).Methods("POST")

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir("./")))

    http.ListenAndServe(":3000", nil)
}