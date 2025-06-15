package handler

import (
	"net/http"

	"github.com/findsam/auth-micro/internal/service"
)

type StoreHandler struct {
	service *service.StoreService
}

func NewStoreHandler(storeHandler *service.StoreService) *StoreHandler {
	return &StoreHandler{
		service: storeHandler,
	}
}

func (h *StoreHandler) Create(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	store, err := h.service.Create(uid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
	}
	SendSuccess(w, r, http.StatusCreated, store)
}
