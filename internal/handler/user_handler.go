package handler

import (
	"net/http"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/service"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		service: userService,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := ParseBody[model.User](r)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	user, tokens, err := h.service.SignUp(user)
	if err != nil {
		SendSuccess(w, r, http.StatusInternalServerError, err)
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

	SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "Successfully signed in",
		"user":    user,
		"token":   tokens.AccessToken,
	})
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	req, err := ParseBody[model.UserSignInRequest](r)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}

	user, tokens, err := h.service.SignIn(req)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
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

	SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "Successfully signed in",
		"user":    user,
		"token":   tokens.AccessToken,
	})
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	user, err := h.service.GetById(uid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "User data retrieved successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := h.service.GetByUsername(username)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "User data retrieved successfully",
		"user":    user,
	})
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	user, err := h.service.GetById(uid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "User data retrieved successfully",
		"user":    user,
	})
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)

	tokens, err := h.service.Refresh(uid)
	if err != nil {
		SendError(w, r, http.StatusInternalServerError, err)
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

	SendSuccess(w, r, http.StatusOK, map[string]any{
		"message": "Successfully refreshed",
		"token":   tokens.AccessToken,
	})
}
