package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "<mansur51>:<ANSKk08aPEDbFjDO>@tcp(127.0.0.1:3306)/<operations>")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}
