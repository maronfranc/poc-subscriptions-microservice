package messages

import (
	"bytes"
	"encoding/json"

	"github.com/maronfranc/poc-subscriptions-microservice/payment/src/rabbitmq"
)

func Transaction(paymentDto PaymentDto) {
	go mockPaymentApproved(paymentDto)
	go mockSubscriptionRequest(paymentDto)
}

func mockPaymentApproved(paymentDto PaymentDto) {
	mock_payment_success := new(bytes.Buffer)
	json.NewEncoder(mock_payment_success).Encode(SuccessDto{
		Message:   "Payment approved",
		Status:    "approved",
		IdProduct: paymentDto.IdProduct,
		IdUser:    paymentDto.IdUser,
	})
	rabbitmq.SendMessage(
		PAYMENTS_E,
		PAYMENTS_APPROVED_Q,
		PAYMENTS_APPROVED_K,
		"topic",
		mock_payment_success.Bytes(),
	)
}

func mockSubscriptionRequest(paymentDto PaymentDto) {
	mock_subscription := new(bytes.Buffer)
	json.NewEncoder(mock_subscription).Encode(ProductSubscriptionDto{
		IdProduct: paymentDto.IdProduct,
		IdUser:    paymentDto.IdUser,
	})
	rabbitmq.SendMessage(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_BUY_REQUEST_Q,
		SUBSCRIPTIONS_BUY_REQUEST_K,
		"topic",
		mock_subscription.Bytes(),
	)
}
