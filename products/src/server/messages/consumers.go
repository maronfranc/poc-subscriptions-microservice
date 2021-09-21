package messages

import (
	"log"

	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
)

// ListenMessageConsumer
func ListenMessageConsumer() {
	log.Println("Listening message broker")

	// Subscription
	go rabbitmq.Consumer(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_BUY_SUCCESS_Q,
		SUBSCRIPTIONS_BUY_SUCCESS_K,
		"topic",
		HandleSubscriptionSuccess,
	)
	go rabbitmq.Consumer(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_BUY_FAIL_Q,
		SUBSCRIPTIONS_BUY_FAIL_K,
		"topic",
		HandleSubscriptionFail,
	)
	// Payment
	go rabbitmq.Consumer(
		PAYMENTS_E,
		PAYMENTS_APPROVED_Q,
		PAYMENTS_APPROVED_K,
		"topic",
		HandlePaymentApproved,
	)
	go rabbitmq.Consumer(
		PAYMENTS_E,
		PAYMENTS_REFUSED_Q,
		PAYMENTS_REFUSED_K,
		"topic",
		HandlePaymentRefused,
	)
}
