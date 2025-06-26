package handler

import (
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
	pid := chi.URLParam(r, "paymentId")
	sid := chi.URLParam(r, "storeId")
	payment, err := h.service.GetById(sid, pid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, payment)
}
