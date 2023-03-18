package gomail

import (
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func Send(from string, to []string, cc []string, subject string, body string, attachment []*multipart.FileHeader) {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Cc", cc...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	for _, attach := range attachment {
		filepath := "/tmp/attachments/" + attach.Filename
		mailer.Attach(filepath)
	}

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

	persist, _ := strconv.ParseBool(os.Getenv("MAIL_ATTACHMENT_PERSIST"))

	if !persist {
		for _, attach := range attachment {
			go DeleteFile(attach)
		}
	}

}

func DeleteFile(attachment *multipart.FileHeader) (err error) {
	filepath := "/tmp/attachments/" + attachment.Filename
	e := os.Remove(filepath)
	if e != nil {
		log.Fatal(e)
	}
	return
}
