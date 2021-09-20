package messages

import (
	"log"

	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
)

// ListenMessageConsumer
func ListenMessageConsumer() {
	log.Println("Listening message broker")

	go rabbitmq.Consumer(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_SUCCESS_Q,
		SUBSCRIPTIONS_BUY_SUCCESS_K,
		"topic",
		HandleSubscriptionSuccess,
	)
	go rabbitmq.Consumer(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_FAIL_Q,
		SUBSCRIPTIONS_BUY_FAIL_K,
		"topic",
		HandleSubscriptionFail,
	)
}
