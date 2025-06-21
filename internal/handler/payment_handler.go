package handler

import (
	"net/http"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/service"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ParseBody[model.CreatePaymentBody](r)
	if err != nil {
		SendError(w, r, http.StatusBadRequest, err)
		return
	}

	payment, err := h.service.Create(body)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	SendSuccess(w, r, http.StatusCreated, payment)
}
