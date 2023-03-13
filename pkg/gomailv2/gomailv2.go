package gomailv2

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func Send(from string, to []string, cc []string, subject string, body string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Cc", cc...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	dialer := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
}
