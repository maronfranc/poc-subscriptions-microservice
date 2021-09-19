package messages

import (
	"log"

	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
)

// SubscriptionBuy send buy message to broker
func SubscriptionBuy(msg []byte) error {
	err := rabbitmq.SendMessage(
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
func HandleDelivery(b []byte) {
	log.Printf(" [x] %s", b)
}
