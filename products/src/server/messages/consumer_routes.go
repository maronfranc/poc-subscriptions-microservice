package messages

import (
	"fmt"

	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
	"github.com/streadway/amqp"
)

// ListenMessageConsumer
func ListenMessageConsumer() {
	go consumer(
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		rabbitmq.SUBSCRIPTIONS_BUY_Q,
		KEY_SUBSCRIPTION_BUY,
		"topic",
		HandleDelivery,
	)
}

func consumer(eName, qName, key, kind string, f func(amqp.Delivery)) error {
	conn, c, err := rabbitmq.ConnectDeclare(eName, qName, key, kind)
	if err != nil {
		return err
	}
	defer c.Close()
	defer conn.Close()

	fmt.Println("Listening message broker:")
	fmt.Println("- Exchange:   ", eName)
	fmt.Println("- Queue:      ", qName)
	fmt.Println("- Routing key:", key)
	rabbitmq.Consume(c, qName, f)

	return nil
}

// Send message to broker
func Send(eName, qName, key, kind string, msg []byte) error {
	conn, c, err := rabbitmq.ConnectDeclare(eName, qName, key, kind)
	if err != nil {
		return err
	}
	defer c.Close()
	defer conn.Close()

	err = rabbitmq.SendMessage(c, eName, key, msg)
	return err
}
