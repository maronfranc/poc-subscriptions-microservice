package messages

import (
	"log"

	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
)

// SubscriptionBuy send buy message to broker
func SubscriptionBuy(msg []byte) error {
	err := rabbitmq.SendMessage(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_REQUEST_Q,
		SUBSCRIPTIONS_BUY_REQUEST_K,
		"topic",
		msg,
	)
	return err
}

// HandleSubscriptionFail receives message from broker
func HandleSubscriptionFail(b []byte) {
	log.Printf(" [x] %s", b)
}

// HandleSubscriptionSuccess receives message from broker
func HandleSubscriptionSuccess(b []byte) {
	log.Println("Handle subscription")
	log.Printf(" [x] %s", b)
}
