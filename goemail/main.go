package main

import (
	"github.com/joho/godotenv"
	"github.com/marceloamoreno87/gomail/pkg/rabbitmq"
)

func main() {
	godotenv.Load("../../.env")
	// for {
	// 	rabbitmq.Publish()
	// }
	rabbitmq.Consume()

}
