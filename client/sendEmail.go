package client

import (
	"log"
	"net/smtp"
)

func SendEmail(email string, body string) {
	from := "...@gmail.com"
	pass := "..."
	to := email

	msg := body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent" + email)
}
