package messages

import (
	"github.com/maronfranc/subscription-system-products/src/rabbitmq"
)

// ListenMessageConsumer
func ListenMessageConsumer() {
	go rabbitmq.Consumer(
		rabbitmq.SUBSCRIPTIONS_BUY_E,
		rabbitmq.SUBSCRIPTIONS_BUY_Q,
		KEY_SUBSCRIPTION_BUY,
		"topic",
		HandleDelivery,
	)
}
