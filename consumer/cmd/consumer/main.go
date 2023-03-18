package main

import (
	"github.com/joho/godotenv"
	"github.com/marceloamoreno87/gomail/consumer/pkg/email"
	"github.com/marceloamoreno87/gomail/consumer/pkg/rabbitmq"
)

func main() {
	godotenv.Load()
	config := rabbitmq.GetConfig()
	rabbitmq.Consume(config, email.Send)
}
