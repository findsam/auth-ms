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

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := ParseRequestBody[model.User](r, h.BaseHandler.Validator)
	if err != nil {
		h.SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	user, tokens, err := h.service.SignUp(user)
	if err != nil {
		h.SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	h.SendSuccess(w, r, http.StatusCreated, map[string]any{
		"user":  user,
		"message": "Signup successful",
		"token": tokens.AccessToken,
	})
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	req, err := ParseRequestBody[model.UserSignInRequest](r, h.BaseHandler.Validator)
	if err != nil {
		h.SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	user, tokens, err := h.service.SignIn(req)
	if err != nil {
		h.SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, 
	})

	h.SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "Successfully signed in",
		"results": map[string]interface{}{
			"token": tokens.AccessToken,
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
