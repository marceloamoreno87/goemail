package rabbitmq

import (
	"fmt"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

type Config struct {
	Queue      string
	Consumer   string
	AutoAck    bool
	Exclusive  bool
	NoLocal    bool
	NoWait     bool
	Durable    bool
	AutoDelete bool
	Args       amqp.Table
}

func GetConfig() *Config {
	autoack, _ := strconv.ParseBool(os.Getenv("AMQP_AUTOACK"))
	exclusive, _ := strconv.ParseBool(os.Getenv("AMQP_EXCLUSIVE"))
	nolocal, _ := strconv.ParseBool(os.Getenv("AMQP_NOLOCAL"))
	nowait, _ := strconv.ParseBool(os.Getenv("AMQP_NOWAIT"))
	durable, _ := strconv.ParseBool(os.Getenv("AMQP_DURABLE"))
	autodelete, _ := strconv.ParseBool(os.Getenv("AMQP_AUTODELETE"))
	config := Config{
		Queue:      os.Getenv("AMQP_QUEUE"),
		Consumer:   os.Getenv("AMQP_CONSUMER"),
		AutoAck:    autoack,
		Exclusive:  exclusive,
		NoLocal:    nolocal,
		NoWait:     nowait,
		Durable:    durable,
		AutoDelete: autodelete,
		Args:       nil,
	}
	return &config
}

func Consume(config *Config, f func(message_body []byte)) {

	connection := getConnection()
	channel := getChannel(connection)
	declareQueue(config, channel)
	msgs := getMessages(config, channel)

	defer connection.Close()
	defer channel.Close()

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			go f(msg.Body)
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
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

func declareQueue(config *Config, channel *amqp.Channel) (queue amqp.Queue, err error) {
	queue, err = channel.QueueDeclare(
		config.Queue,
		config.Durable,
		config.AutoDelete,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		panic(err)
	}

	return
}
