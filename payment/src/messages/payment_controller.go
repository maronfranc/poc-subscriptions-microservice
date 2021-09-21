package messages

import (
	"encoding/json"
)

type PaymentDto struct {
	PaymentMethod string  `json:"payment_method"`
	Price         float32 `json:"amount"`
	IdProduct     string  `json:"id_product"`
	IdUser        string  `json:"id_user"`
}

type SuccessDto struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	IdProduct string `json:"id_product"`
	IdUser    string `json:"id_user"`
}

type ProductSubscriptionDto struct {
	IdProduct string `json:"id_product"`
	IdUser    string `json:"id_user"`
}

func handleTransactionRequest(b []byte) {
	var p PaymentDto
	if err := json.Unmarshal(b, &p); err != nil {
		// TODO: dead-letter
		panic(err)
	}
	Transaction(p)
}
