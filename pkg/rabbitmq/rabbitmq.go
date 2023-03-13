package rabbitmq

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/marceloamoreno87/gomail/pkg/email"
	"github.com/streadway/amqp"
)

type Config struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func GetConfig() *Config {
	autoack, _ := strconv.ParseBool(os.Getenv("AMQP_AUTOACK"))
	exclusive, _ := strconv.ParseBool(os.Getenv("AMQP_EXCLUSIVE"))
	nolocal, _ := strconv.ParseBool(os.Getenv("AMQP_NOLOCAL"))
	nowait, _ := strconv.ParseBool(os.Getenv("AMQP_NOWAIT"))
	config := Config{
		Queue:     os.Getenv("AMQP_QUEUE"),
		Consumer:  os.Getenv("AMQP_CONSUMER"),
		AutoAck:   autoack,
		Exclusive: exclusive,
		NoLocal:   nolocal,
		NoWait:    nowait,
		Args:      nil,
	}
	return &config
}

func Consume(config *Config, f func(message_body []byte)) {

	connection := getConnection()
	channel := getChannel(connection)
	msgs := getMessages(config, channel)

	defer connection.Close()
	defer channel.Close()

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			f(msg.Body)
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}

func Publish() {

	mock := email.MailMessage{
		Params:  []string{"1", "2", "3"},
		To:      []string{"to@to.com.br", "robin@robin.com.br"},
		Cc:      []string{"cc@cc.com.br", "batman@batman.com.br"},
		Subject: "Testando",
		From:    "lal@teste.com.br",
		Body:    "lalala",
	}

	json_data, err := json.Marshal(mock)

	connection := getConnection()
	channel := getChannel(connection)

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

func getConnection() (connection *amqp.Connection) {
	connection, err := amqp.Dial(os.Getenv("AMQP_SERVER_URL"))
	if err != nil {
		panic(err)
	}
	return
}

func getChannel(connection *amqp.Connection) (channel *amqp.Channel) {
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return
}

func getMessages(config *Config, channel *amqp.Channel) (msgs <-chan amqp.Delivery) {
	msgs, err := channel.Consume(
		config.Queue,
		config.Consumer,
		config.AutoAck,
		config.Exclusive,
		config.NoLocal,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		panic(err)
	}
	return
}
