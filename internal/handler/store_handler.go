package handler

import "net/http"


type StoreHandler struct {}

func NewStoreHandler() *StoreHandler {
	return &StoreHandler{}
}

func (h *StoreHandler) GetStore(w http.ResponseWriter, r *http.Request) {}