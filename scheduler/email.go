package main

import (
	"log"
	"net/smtp"
)

func sendMessage(message []byte) {
	email := "user@gmail.com"

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		email,
		"********",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		[]string{email},
		message,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sendMessage([]byte("This is the email body."))
}
