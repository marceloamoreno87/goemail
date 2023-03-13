package main

import (
	"github.com/joho/godotenv"
	"github.com/marceloamoreno87/gomail/pkg/email"
	"github.com/marceloamoreno87/gomail/pkg/rabbitmq"
)

func main() {
	godotenv.Load("../../.env")
	// publish()
	config := rabbitmq.GetConfig()
	rabbitmq.Consume(config, email.Send)
}

// Just for mocking
func publish() {
	for {
		rabbitmq.Publish()
	}
}
