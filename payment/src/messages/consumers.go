package messages

import "github.com/maronfranc/poc-subscriptions-microservice/payment/src/rabbitmq"

// ListenMessageConsumer
func ListenMessageConsumer() {
	go rabbitmq.Consumer(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_BUY_REQUEST_Q,
		SUBSCRIPTIONS_BUY_REQUEST_K,
		"topic",
		handleTransactionRequest,
	)
}
