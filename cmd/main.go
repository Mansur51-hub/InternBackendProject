package main

import (
	Models "Microservice/Models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"net/http"
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
	stmt, err := db.Prepare("INSERT INTO users(id, balance) VALUES(?, ?)")
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

	if !userExists(id) {
		fmt.Println("add user")
		_, err = stmt.Exec(id, balance)

		if err != nil {
			panic(err.Error())
		}
	} else {
		_, err = db.Query("update users set balance = balance + ? where id = ?", balance, id)

		if err != nil {
			panic(err.Error())
		}
	}

	stmt.Close()

	fmt.Fprintf(w, "New user was addedd")
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

	if userExists(id) {
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

	} else {
		fmt.Fprintf(w, "User was not found")
		json.NewEncoder(w).Encode(Models.User{})
	}

}

func userExists(id string) bool {
	resp, err := db.Query("select * from users where id = ?", id)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		fmt.Println("Db does not response correctly")
		return false
	}
	defer resp.Close()

	return resp.Next()
}
