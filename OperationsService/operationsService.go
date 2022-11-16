package OperationsService

import (
	"Microservice/Models"
	"Microservice/Repository"
)

type OperationService interface {
	TopUpMoney(id int, balance float32) error
	ReserveMoney(orderId int, userId int, serviceId int, amount float32) error
	WriteOffFromTheReservedMoney(orderId int, userId int, serviceId int, amount float32) error
	GetMoneyAmount(id int) (Models.User, error)
	IsUserExists(id int) (bool, error)
}

func NewOperationService(repository *Repository.RepositoryMySql) *MyOperationsService {
	return NewMyOperationsService(repository)
}
