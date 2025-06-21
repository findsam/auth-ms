package handler

import (
	"fmt"
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
	body, err := ParseBody[model.Payment](r)
	if err != nil {
		SendError(w, r, http.StatusBadRequest, err)
		return
	}
	payment, err := h.service.Create(body)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("%+v\n", payment)
	SendSuccess(w, r, http.StatusCreated, payment)
}