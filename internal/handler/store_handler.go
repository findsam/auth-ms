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

func (h *StoreHandler) GetStore(w http.ResponseWriter, r *http.Request) {}