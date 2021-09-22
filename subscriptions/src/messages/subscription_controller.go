package messages

import "encoding/json"

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

func HandleSubscriptionRequest(b []byte) {
	var dto ProductSubscriptionDto
	if err := json.Unmarshal(b, &dto); err != nil {
		panic(err)
	}
	CreateSubscription(dto)
}
