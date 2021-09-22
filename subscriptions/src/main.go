package main

import (
	"github.com/maronfranc/poc-subscriptions-microservice/subscriptions/src/messages"
)

func main() {
	ch := make(chan interface{})
	messages.ListenMessageConsumer()
	<-ch
}
