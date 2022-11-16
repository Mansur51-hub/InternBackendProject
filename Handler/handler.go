package Handler

import (
	"Microservice/OperationsService"
)

type Handler struct {
	service *OperationsService.MyOperationsService
}

func NewHandler(service *OperationsService.MyOperationsService) *Handler {
	return &Handler{service}
}
