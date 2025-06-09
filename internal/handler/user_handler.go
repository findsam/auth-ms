package handler

import (
	"encoding/json"
	"net/http"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/service"
	"github.com/findsam/auth-micro/pkg/util"
	"github.com/go-chi/render"
)

type UserHandler struct {
	service   *service.UserService
	validator *util.Validator
}

func NewUserHandler(userService *service.UserService, validator *util.Validator) *UserHandler {
	return &UserHandler{
		service:   userService,
		validator: validator,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]any{
			"messages": h.validator.ParseValidationErrors(err),
		})
		return
	}

	user, err := h.service.CreateUser(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]any{
			"message": "Failed to create user",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]any{
		"result": user,
		"messages": "User created successfully",
	})

}
