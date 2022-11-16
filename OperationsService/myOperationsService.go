package OperationsService

import (
	"Microservice/Models"
	"Microservice/Repository"
	"errors"
)

type MyOperationsService struct {
	repository *Repository.RepositoryMySql
}

func NewMyOperationsService(repository *Repository.RepositoryMySql) *MyOperationsService {
	return &MyOperationsService{repository}
}

func (myService *MyOperationsService) TopUpMoney(id int, balance float32) (Models.User, error) {
	return myService.repository.TopUpBalance(id, balance)
}

func (myService *MyOperationsService) ReserveMoney(orderId int, userId int, serviceId int, amount float32) error {
	rep := myService.repository

	if _, err := rep.UserExists(userId); err != nil {
		return errors.New("user not found")
	}

	if _, err := rep.OrderReserved(orderId); err != nil {
		return errors.New("order is already reserved")
	}

	balance, err := rep.GetBalance(userId)
	if err != nil {
		return err
	}

	if balance.Amount-amount < 0 {
		return errors.New("no money enough")
	}

	err = rep.ReserveAmount(orderId, userId, serviceId, amount)

	if err != nil {
		return err
	}

	return nil
}

func (myService *MyOperationsService) WriteOffFromTheReservedMoney(orderId int, userId int, serviceId int, amount float32) error {
	rep := myService.repository

	reserved, err := rep.OrderReserved(orderId)

	if err != nil {
		return err
	}

	if !reserved {
		return errors.New("no such reservation found")
	}

	err = rep.WriteOffFromTheReserveMoney(orderId)

	if err != nil {
		return errors.New("could not write off money")
	}

	err = rep.AddToReport(orderId, userId, serviceId, amount)

	if err != nil {
		return errors.New("could add row to report")
	}

	return nil
}

func (myService *MyOperationsService) GetMoneyAmount(id int) (Models.User, error) {
	return myService.repository.GetBalance(id)
}

func (myService *MyOperationsService) IsUserExists(id int) (bool, error) {
	return false, nil
}
