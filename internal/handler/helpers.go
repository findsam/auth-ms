package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/findsam/auth-micro/pkg/util"
	"github.com/go-chi/render"
)

var validator = util.NewValidator()

func SendError(w http.ResponseWriter, r *http.Request, status int, err error) {
	render.Status(r, status)
	render.JSON(w, r, map[string]any{
		"message": err.Error(),
	})
}

func SendSuccess(w http.ResponseWriter, r *http.Request, status int, data any) {
	render.Status(r, status)
	render.JSON(w, r, data)
}

func ParseBody[T any](r *http.Request) (*T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid JSON body")
	}
	defer r.Body.Close()

	if err := validator.Struct(data); err != nil {
		return nil, err
	}

	return &data, nil
}
