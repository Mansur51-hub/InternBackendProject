package Handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) TopUpBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	balance := keyVal["balance"]

	idInt, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	balanceFloat, err := strconv.ParseFloat(balance, 32)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	err = h.service.TopUpMoney(idInt, float32(balanceFloat))

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]

	idInt, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error input values")
	}

	user, err := h.service.GetMoneyAmount(idInt)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}

	json.NewEncoder(w).Encode(user)

}
