package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "1510819@hcmut.edu.vn", "AnhY3u3m410", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"tranvanduc.mt151@gmail.com"}
	msg := []byte("To: 1510819@hcmut.edu.vn\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "1510819@hcmut.edu.vn", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
