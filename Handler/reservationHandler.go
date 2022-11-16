package Handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) ReserveMoney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	orderId := keyVal["order_id"]
	userId := keyVal["user_id"]
	serviceId := keyVal["service_id"]
	amount := keyVal["amount"]

	orderIdInt, err := strconv.Atoi(orderId)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	serviceIdInt, err := strconv.Atoi(serviceId)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	amountFloat, err := strconv.ParseFloat(amount, 32)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	err = h.service.ReserveMoney(orderIdInt, userIdInt, serviceIdInt, float32(amountFloat))

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}
}

func (h *Handler) WriteOffFromTheReservedMoney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	orderId := keyVal["order_id"]
	userId := keyVal["user_id"]
	serviceId := keyVal["service_id"]
	amount := keyVal["amount"]

	orderIdInt, err := strconv.Atoi(orderId)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	serviceIdInt, err := strconv.Atoi(serviceId)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	amountFloat, err := strconv.ParseFloat(amount, 32)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	err = h.service.WriteOffFromTheReservedMoney(orderIdInt, userIdInt, serviceIdInt, float32(amountFloat))

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}
}
