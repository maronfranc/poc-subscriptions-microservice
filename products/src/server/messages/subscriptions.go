package messages

import (
	"log"

	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
	"github.com/streadway/amqp"
)

// SubscriptionBuy send buy message to broker
func SubscriptionBuy(msg []byte) error {
	err := Send(
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		rabbitmq.SUBSCRIPTIONS_BUY_Q,
		KEY_SUBSCRIPTION_BUY,
		"topic",
		msg,
	)
	if err != nil {
		return err
	}
	return err
}

// HandleDelivery receives message from broker
func HandleDelivery(d amqp.Delivery) {
	log.Printf(" [x] %s", d.Body)
}
