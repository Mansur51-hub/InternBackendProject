package Repository

import (
	"database/sql"
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

	exists, err := repository.userExists(id)

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

func (repository *RepositoryMySql) reserveAmount(orderId int, userId int, serviceId int, amount float32) error {
	return nil
}

func (repository *RepositoryMySql) writeOffFromTheReserveMoney(orderId int, userId int, serviceId int, amount float32) error {
	return nil
}

func (repository *RepositoryMySql) getBalance(id int) error {
	return nil
}

func (repository *RepositoryMySql) userExists(id int) (bool, error) {
	resp, err := repository.db.Query("select * from users where id = ?", id)
	if err != nil {
		return false, err
	}
	defer resp.Close()

	return resp.Next(), nil
}
