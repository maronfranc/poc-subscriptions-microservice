package rabbitmq

import (
	"fmt"
	"log"

	"github.com/maronfranc/subscription-system-products/src/config"
	"github.com/streadway/amqp"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

// Connect to amqp keep it open and return open channel
func Connect() (*amqp.Connection, *amqp.Channel, error) {
	cfg := config.Cfg.Rabbitmq
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, nil, err
	}

	c, err := conn.Channel()
	if err != nil {
		return conn, c, err
	}

	return conn, c, err
}

func DeclareExchange(c *amqp.Channel, eName string, kind string) error {
	err := c.ExchangeDeclare(
		eName, // name
		kind,  // type
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	return err
}

func QueueDeclare(c *amqp.Channel, queueName string) (*amqp.Queue, error) {
	q, err := c.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	return &q, err
}

func BindQueueToExchange(c *amqp.Channel, q *amqp.Queue, exchangeName, key string) error {
	err := c.QueueBind(
		q.Name,       // queue name
		key,          // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	return err
}

// SendMessage to exchange
func SendMessage(c *amqp.Channel, eName, rountingKey string, msg []byte) error {
	if err := c.Publish(
		eName,       // exchange
		rountingKey, // queue name
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		}, // message to publish
	); err != nil {
		return err
	}

	return nil
}

// Consumer create exchange and queue connection and handle delivery data
func Consumer(eName, qName, key, kind string, f func(amqp.Delivery)) error {
	conn, c, err := ConnectDeclare(eName, qName, key, kind)
	if err != nil {
		return err
	}
	defer c.Close()
	defer conn.Close()

	log.Println("Listening message broker:")
	log.Println("- Exchange:   ", eName)
	log.Println("- Queue:      ", qName)
	log.Println("- Routing key:", key)
	consume(c, qName, f)

	return nil
}

// ConnectDeclare create queue and bind to exchange
func ConnectDeclare(eName, qName, key, kind string) (*amqp.Connection, *amqp.Channel, error) {
	conn, c, err := Connect()
	if err != nil {
		return conn, c, err
	}

	if err := DeclareExchange(c, eName, kind); err != nil {
		return conn, c, err
	}

	q, err := QueueDeclare(c, qName)
	if err != nil {
		return conn, c, err
	}

	BindQueueToExchange(c, q, eName, key)
	return conn, c, err
}

// Send message to broker
func Send(eName, qName, key, kind string, msg []byte) error {
	conn, c, err := ConnectDeclare(eName, qName, key, kind)
	if err != nil {
		return err
	}
	defer c.Close()
	defer conn.Close()

	err = SendMessage(c, eName, key, msg)
	return err
}

func consume(c *amqp.Channel, qName string, f func(amqp.Delivery)) {
	consumer, err := c.Consume(
		qName, // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range consumer {
			f(d)
		}
	}()
	<-forever
}
