package gomail

import (
	"log"
	"os"
	"strconv"

	"github.com/marceloamoreno87/gomail/consumer/pkg/email"

	gomail "gopkg.in/gomail.v2"
)

func Send(mailmessage *email.MailMessage) {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", mailmessage.GetFrom())
	mailer.SetHeader("To", mailmessage.GetTo()...)
	mailer.SetHeader("Cc", mailmessage.GetCc()...)
	mailer.SetHeader("Subject", mailmessage.GetSubject())
	mailer.SetBody("text/html", mailmessage.GetBody())

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
