package messages

import (
	"log"

	"github.com/maronfranc/poc-subscriptions-microservice/subscriptions/src/rabbitmq"
)

// ListenMessageConsumer
func ListenMessageConsumer() {
	log.Println("Listening message broker")

	// Subscription
	go rabbitmq.Consumer(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_BUY_REQUEST_Q,
		SUBSCRIPTIONS_BUY_REQUEST_K,
		"topic",
		HandleSubscriptionRequest,
	)
}
