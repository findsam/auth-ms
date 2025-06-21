package handler

import "github.com/findsam/auth-micro/internal/service"
type PaymentHandler struct {
	service *service.PaymentService
}
func NewPaymentHandler(service *service.StoreService) *StoreHandler {
	return &StoreHandler{
		service: service,
	}
}
