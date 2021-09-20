package main

import (
	"github.com/maronfranc/poc-subscriptions-microservice/payment/src/config"
	"github.com/maronfranc/poc-subscriptions-microservice/payment/src/messages"
)

func main() {
	config.Cfg = config.GetConfig()
	messages.ListenMessageConsumer()
}
