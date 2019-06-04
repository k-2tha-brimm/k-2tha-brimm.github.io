// package main

// import (
// 	// "strings"
// 	"fmt"
// 	"net/smtp"
// )

// type Message struct {
// 	Email string
// 	Content string
// 	Errors map[string]string
// }

// func (msg *Message) Deliver() error {
// 	to := []string{"someone@example.com"}
// 	body := fmt.Sprintf("reply-To: %v\r\nSubject: New Message\r\n%v", msg.Email, msg.Content)

// 	username := "kev.brimm@gmail.com"
// 	password := "Stanf0rdsux!"
// 	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

// 	return smtp.SendMail("smtp.gmail.com:587", auth, msg.Email, to, []byte(body))
// }