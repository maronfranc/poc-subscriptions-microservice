package messages

import (
	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
	"github.com/streadway/amqp"
)

// SubscriptionBuy send buy message to broker
func SubscriptionBuy(msg []byte) error {
	conn, c, err := connectDeclare(
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		rabbitmq.SUBSCRIPTIONS_BUY_Q,
		KEY_SUBSCRIPTION_BUY,
	)
	if err != nil {
		return err
	}
	defer c.Close()
	defer conn.Close()

	err = rabbitmq.SendMessage(
		c,
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		KEY_SUBSCRIPTION_BUY,
		msg,
	)
	return err
}

// SubscriptionSuccess receive message from broker
func SubscriptionSuccess() error {
	conn, c, err := connectDeclare(
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		rabbitmq.SUBSCRIPTIONS_BUY_Q,
		KEY_SUBSCRIPTION_BUY,
	)
	if err != nil {
		return err
	}
	defer c.Close()
	defer conn.Close()

	rabbitmq.Consume(c, rabbitmq.SUBSCRIPTIONS_BUY_Q)

	return err
}

func connectDeclare(eName, qName, key string) (*amqp.Connection, *amqp.Channel, error) {
	conn, c, err := rabbitmq.Connect()
	if err != nil {
		return conn, c, err
	}

	if err := rabbitmq.DeclareExchange(c, eName, "topic"); err != nil {
		return conn, c, err
	}

	q, err := rabbitmq.QueueDeclare(c, qName)
	if err != nil {
		return conn, c, err
	}

	rabbitmq.BindQueueToExchange(c, q, eName, key)
	return conn, c, err
}
