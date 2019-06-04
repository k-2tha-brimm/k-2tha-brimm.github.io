package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"net/smtp"
)

func send(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var message Message
	// _ = json.NewDecoder(r.Body).Decode(&message)
	msg := &Message{
		Email: r.FormValue("email"),
		Content: r.FormValue("content"),
	}

	if err := msg.Deliver(); err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", send).Methods("POST")

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir("./")))

    http.ListenAndServe(":3000", nil)
}

type Message struct {
	Email string
	Content string
	Errors map[string]string
}

func (msg *Message) Deliver() error {
	to := []string{"someone@example.com"}
	body := fmt.Sprintf("reply-To: %v\r\nSubject: New Message\r\n%v", msg.Email, msg.Content)

	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	return smtp.SendMail("smtp.gmail.com:587", auth, msg.Email, to, []byte(body))
}