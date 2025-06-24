package handler

import (
	"net/http"

	"github.com/findsam/auth-micro/internal/service"
	"github.com/go-chi/chi/v5"
)

type StoreHandler struct {
	service *service.StoreService
}

func NewStoreHandler(service *service.StoreService) *StoreHandler {
	return &StoreHandler{
		service: service,
	}
}

func (h *StoreHandler) Create(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	store, err := h.service.Create(uid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusCreated, store)
}

func (h *StoreHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	store, err := h.service.GetById(id)

	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, store)
}
