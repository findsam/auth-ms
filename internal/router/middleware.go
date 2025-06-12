package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/pkg/token"
	"github.com/golang-jwt/jwt/v5"
)

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if len(tokenAuth) > 7 && tokenAuth[:7] == "Bearer " {
		return tokenAuth[7:]
	}
	return ""
}

func WithJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		str := getTokenFromRequest(r)
		fmt.Println(str)

		t, err := token.ValidateJWT(str)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				if r.URL.Path == "/api/v1/users/refresh" {
					ctx := context.WithValue(r.Context(), "uid", token.ReadJWT(t))
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				handler.SendError(w, r, http.StatusUnauthorized, fmt.Errorf("expired"))
				return
			}
			handler.SendError(w, r, http.StatusUnauthorized, err)
			return
		}

		if !t.Valid {
			handler.SendError(w, r, http.StatusUnauthorized, fmt.Errorf("token is not valid"))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", token.ReadJWT(t))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
