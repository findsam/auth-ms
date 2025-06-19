package handler

import (
	"net/http"
	"time"

	"github.com/findsam/auth-micro/internal/service"
	"github.com/go-chi/chi/v5"
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

func (h *StoreHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	store, err := h.service.GetById(id)
	time.Sleep(3 * time.Second)

	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, store)
}
