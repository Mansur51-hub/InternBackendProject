package Repository

import (
	"Microservice/Models"
	"database/sql"
	"errors"
	"time"
)

type RepositoryMySql struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) *RepositoryMySql {
	return &RepositoryMySql{db}
}

func (repository *RepositoryMySql) TopUpBalance(id int, balance float32) error {
	stmt, err := repository.db.Prepare("INSERT INTO users(id, balance) VALUES(?, ?)")
	if err != nil {
		return err
	}

	exists, err := repository.UserExists(id)

	if err != nil {
		return err
	}

	if !exists {
		_, err = stmt.Exec(id, balance)

		if err != nil {
			return err
		}
	} else {
		_, err = repository.db.Query("update users set balance = balance + ? where id = ?", balance, id)

		if err != nil {
			return err
		}
	}

	stmt.Close()

	return nil
}

func (repository *RepositoryMySql) ReserveAmount(orderId int, userId int, serviceId int, amount float32) error {
	_, err := repository.db.Query("update users set balance = balance - ? where id = ?", amount, userId)

	if err != nil {
		return err
	}

	_, err = repository.db.Query("insert into reservations (order_id, user_id, service_id, amount) VALUES(?, ?, ?, ?)", orderId, userId, serviceId, amount)

	if err != nil {
		repository.db.Query("update users set balance = balance + ? where id = ?", amount, userId)
		return err
	}

	return nil
}

func (repository *RepositoryMySql) WriteOffFromTheReserveMoney(orderId int) error {
	_, err := repository.db.Query("delete from reservations where order_id = ?", orderId)

	if err != nil {
		return err
	}

	return nil
}

func (repository *RepositoryMySql) GetBalance(id int) (Models.User, error) {
	userExists, err := repository.UserExists(id)

	if err != nil {
		return Models.User{}, err
	}

	if userExists {
		resp, err := repository.db.Query("select * from users where id = ?", id)

		if err != nil {
			return Models.User{}, err
		}

		defer resp.Close()

		var user Models.User

		for resp.Next() {
			err := resp.Scan(&user.Id, &user.Amount)
			if err != nil {
				return Models.User{}, err
			}
		}

		return user, err
	} else {
		return Models.User{}, errors.New("user not found")
	}
}

func (repository *RepositoryMySql) UserExists(id int) (bool, error) {
	resp, err := repository.db.Query("select * from users where id = ?", id)
	if err != nil {
		return false, err
	}
	defer resp.Close()

	return resp.Next(), nil
}

func (repository *RepositoryMySql) OrderReserved(id int) (bool, error) {
	resp, err := repository.db.Query("select order_id from reservations where order_id = ?", id)

	if err != nil {
		return false, err
	}
	defer resp.Close()

	return resp.Next(), nil
}

func (repository *RepositoryMySql) OrderExist(id int) (bool, error) {
	resp, err := repository.db.Query("select * from users where order_id = ?", id)
	if err != nil {
		return false, err
	}
	defer resp.Close()

	return resp.Next(), nil
}

func (repository *RepositoryMySql) AddToReport(orderId int, userId int, serviceId int, amount float32) error {
	resp, err := repository.db.Query("INSERT INTO reports(order_id, user_id, service_id, amount, date) VALUES(?, ?, ?, ?, ?)", orderId, userId, serviceId, amount, time.Now().Format("2006-01-02"))
	defer resp.Close()

	return err
}
