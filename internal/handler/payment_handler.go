package handler

import (
	"fmt"
	"net/http"

	"github.com/findsam/auth-micro/internal/service"
	"github.com/go-chi/chi/v5"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) GetById(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	paymentId := chi.URLParam(r, "paymentId")
	fmt.Println(username, paymentId)
	result, err := h.service.GetById(username, paymentId)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, result) 
}
