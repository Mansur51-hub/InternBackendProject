package Repository

import "database/sql"

type Repository interface {
	topUpBalance(id int, balance float32) error
	reserveAmount(orderId int, userId int, serviceId int, amount float32) error
	writeOffFromTheReserveMoney(orderId int, userId int, serviceId int, amount float32) error
	getBalance(id int) error
	userExists(id int) (bool, error)
}

func NewRepository(db *sql.DB) *RepositoryMySql {
	return NewRepositoryMySql(db)
}
