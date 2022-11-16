package main

import (
	"Microservice/Models"
	"Microservice/Repository"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:1111@tcp(127.0.0.1:3306)/operations")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/balances", topUpBalance).Methods("POST")
	router.HandleFunc("/balances", getBalance).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func topUpBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		panic(err.Error())
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	balance := keyVal["balance"]

	myrep := Repository.NewRepository(db)

	idInt, err := strconv.Atoi(id)
	balanceFloat, err := strconv.ParseFloat(balance, 32)
	myrep.TopUpBalance(idInt, float32(balanceFloat))

}

func getBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("2")

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]

	fmt.Println("3")

	resp, err := db.Query("select * from users where id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Close()

	var user Models.User

	for resp.Next() {
		err := resp.Scan(&user.Id, &user.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(user)

}
