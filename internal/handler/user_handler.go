package handler

import (
	"net/http"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/service"
)

type UserHandler struct {
	*BaseHandler
	service *service.UserService
}

func NewUserHandler(baseHandler *BaseHandler, userService *service.UserService) *UserHandler {
	return &UserHandler{
		BaseHandler: baseHandler,
		service:     userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := ParseRequestBody[model.User](r, h.BaseHandler.Validator)
	if err != nil {
		h.SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	user, tokens, err := h.service.CreateUser(user)
	if err != nil {
		h.SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: tokens.RefreshToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:  true,
		Path: "/",
	})

	h.SendSuccess(w, r, http.StatusCreated, "User Created Successfully", map[string]any{
		"user":        user,
		"token": tokens.AccessToken,
	})
}
