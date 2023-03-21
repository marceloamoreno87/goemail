package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/marceloamoreno87/gomail/consumer/pkg/driver/gomail"
	"github.com/marceloamoreno87/gomail/consumer/pkg/driver/sendgrid"
	"github.com/marceloamoreno87/gomail/consumer/pkg/driver/ses"
	"github.com/marceloamoreno87/gomail/consumer/pkg/email"
	"github.com/marceloamoreno87/gomail/consumer/pkg/rabbitmq"
)

func main() {
	godotenv.Load("../../.env")
	config := rabbitmq.GetConfig()
	rabbitmq.Consume(config, send)
}

func send(message_body []byte) {
	mailmessage := email.SetMailMessage(message_body)
	switch os.Getenv("MAIL_DRIVER") {
	case "gomail":
		gomail.Send(mailmessage)
	case "sendgrid":
		sendgrid.Send(mailmessage)
	case "ses":
		ses.Send(mailmessage)
	default:
		gomail.Send(mailmessage)
	}
}
