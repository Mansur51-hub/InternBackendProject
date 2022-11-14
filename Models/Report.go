package Models

import "time"

type T struct {
	OrderId   string    `json:"order_id"`
	ServiceId string    `json:"service_id"`
	UserId    string    `json:"user_id"`
	Amount    float32   `json:"amount"`
	Date      time.Time `json:"date"`
}
