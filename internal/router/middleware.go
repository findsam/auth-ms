package router

import (
	"context"
	"net/http"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/pkg/token"
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
		t, err := token.ValidateJWT(str)

		if err != nil || !t.Valid {
			handler.SendError(w, r, http.StatusInternalServerError, err)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", token.ReadJWT(t))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}