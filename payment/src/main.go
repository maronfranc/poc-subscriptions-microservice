package main

import (
	"github.com/maronfranc/poc-subscriptions-microservice/payment/src/config"
	"github.com/maronfranc/poc-subscriptions-microservice/payment/src/messages"
)

func main() {
	// load configuration to be used in rabbitmq connection
	config.Cfg = config.GetConfig()

	ch := make(chan interface{})
	messages.ListenMessageConsumer()
	<-ch
}
