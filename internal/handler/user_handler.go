package handler

import (
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
	user, err := ParseRequestBody[model.User](r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "error")
		return
	}
	err = ValidateBody(user, h.validator)
	if err != nil { 
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]any{
			"messages": h.validator.ParseValidationErrors(err),
		})
		return
	}

	if err != nil { 
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]any{
			"messages": h.validator.ParseValidationErrors(err),
		})
		return
	}



	// if err := h.validator.Struct(user); err != nil {
	// 	render.Status(r, http.StatusBadRequest)
	// 	render.JSON(w, r, map[string]any{
	// 		"messages": h.validator.ParseValidationErrors(err),
	// 	})
	// 	return
	// }

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]any{
			"message": "Failed to create user",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]any{
		"result":   createdUser,
		"messages": "User created successfully",
	})
}
