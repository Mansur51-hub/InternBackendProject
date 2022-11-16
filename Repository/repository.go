package Repository

import (
	"Microservice/Models"
	"database/sql"
)

type Repository interface {
	TopUpBalance(id int, balance float32) (Models.User, error)
	ReserveAmount(orderId int, userId int, serviceId int, amount float32) error
	WriteOffFromTheReserveMoney(orderId int) error
	GetBalance(id int) (Models.User, error)
	UserExists(id int) (bool, error)
	OrderReserved(id int) (bool, error)
	OrderExists(id int) (bool, error)
	AddToReport(orderId int, userId int, serviceId int, amount float32) error
}

func NewRepository(db *sql.DB) *RepositoryMySql {
	return NewRepositoryMySql(db)
}
