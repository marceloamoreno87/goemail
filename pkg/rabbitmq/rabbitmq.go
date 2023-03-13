package rabbitmq

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marceloamoreno87/gomail/pkg/email"
	"github.com/streadway/amqp"
)

func Consume() {

	connection := getConnection()
	channel := getChannel(connection)
	msgs := getMessages(channel)

	defer connection.Close()
	defer channel.Close()

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			email.Send(msg.Body)
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}

func getConnection() *amqp.Connection {
	connection, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		panic(err)
	}
	return connection
}

func getChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return channel
}
func getMessages(channel *amqp.Channel) <-chan amqp.Delivery {
	msgs, err := channel.Consume(
		"mail_messages",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return msgs
}

func Publish() {

	teste := email.MailMessage{
		Params:  []string{"1", "2", "3"},
		To:      []string{"to@to.com.br", "robin@robin.com.br"},
		Cc:      []string{"cc@cc.com.br", "batman@batman.com.br"},
		Subject: "Testando",
		From:    "lal@teste.com.br",
		Body:    "lalala",
	}

	json_data, err := json.Marshal(teste)

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"mail_messages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = channel.Publish(
		"",
		"mail_messages",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        json_data,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
}
