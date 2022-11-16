package main

import (
	"Microservice/Handler"
	"Microservice/OperationsService"
	"Microservice/Repository"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:1111@tcp(127.0.0.1:3306)/operations")
	if err != nil {
		panic(err.Error())
	}
	rep := Repository.NewRepository(db)
	service := OperationsService.NewOperationService(rep)
	handler := Handler.NewHandler(service)

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/balances", handler.TopUpBalance).Methods("POST")
	router.HandleFunc("/balances", handler.GetBalance).Methods("GET")
	router.HandleFunc("/reservations", handler.ReserveMoney).Methods("POST")
	router.HandleFunc("/reservations", handler.WriteOffFromTheReservedMoney).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
