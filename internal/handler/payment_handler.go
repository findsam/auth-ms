package handler

import (
	"net/http"

	"github.com/findsam/auth-micro/internal/model"
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

func (h *PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ParseBody[model.CreatePaymentBody](r)
	if err != nil {
		SendError(w, r, http.StatusBadRequest, err)
		return
	}

	payment, err := h.service.Create(body.StoreId)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	SendSuccess(w, r, http.StatusCreated, payment)
}

func (h *PaymentHandler) GetByStoreId(w http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "id")
	payments, err := h.service.GetByStoreId(sid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, payments)

}