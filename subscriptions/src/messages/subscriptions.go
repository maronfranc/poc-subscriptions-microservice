package messages

import (
	"bytes"
	"encoding/json"

	"github.com/maronfranc/poc-subscriptions-microservice/subscriptions/src/rabbitmq"
)

func CreateSubscription(dto ProductSubscriptionDto) {
	go mockSubscriptionSuccess(dto)
}

func mockSubscriptionSuccess(dto ProductSubscriptionDto) {
	mock_subscription_success := new(bytes.Buffer)
	json.NewEncoder(mock_subscription_success).Encode(SuccessDto{
		Message:   "Subscription approved",
		Status:    "approved",
		IdProduct: dto.IdProduct,
		IdUser:    dto.IdUser,
	})
	rabbitmq.SendMessage(
		SUBSCRIPTIONS_E,
		SUBSCRIPTIONS_BUY_SUCCESS_Q,
		SUBSCRIPTIONS_BUY_SUCCESS_K,
		"topic",
		mock_subscription_success.Bytes(),
	)
}
