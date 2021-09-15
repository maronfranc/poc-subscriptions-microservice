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

func Consume(c *amqp.Channel, qName string) {
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

	listen := make(chan bool)

	go func() {
		for d := range consumer {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	<-listen
}
