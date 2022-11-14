package Models

type Reserving struct {
	OrderId   string  `json:"order_id"`
	UserId    string  `json:"user_id"`
	ServiceId string  `json:"service_id"`
	Amount    float32 `json:"amount"`
}
