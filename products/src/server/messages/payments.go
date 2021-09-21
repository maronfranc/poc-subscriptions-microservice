package messages

import "log"

func HandlePaymentApproved(b []byte) {
	log.Println("HandlePaymentApproved")
	log.Printf(" [x] %s", b)
}

func HandlePaymentRefused(b []byte) {
	log.Println("HandlePaymentRefused")
	log.Printf(" [x] %s", b)
}
