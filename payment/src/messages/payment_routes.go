package messages

import (
	"github.com/maronfranc/poc-subscriptions-microservice/payment/src/rabbitmq"
)

// ListenMessageConsumer
func ListenMessageConsumer() {
	rabbitmq.Consumer(
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		rabbitmq.SUBSCRIPTIONS_BUY_Q,
		KEY_SUBSCRIPTION_BUY,
		"topic",
		handleTransactionRequest,
	)
}
