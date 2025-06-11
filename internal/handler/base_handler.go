package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/findsam/auth-micro/pkg/util"
	"github.com/go-chi/render"
)

type BaseHandler struct {
	Validator *util.Validator
}

func NewBaseHandler(validator *util.Validator) *BaseHandler {
	return &BaseHandler{
		Validator: validator,
	}
}

func (h *BaseHandler) SendError(w http.ResponseWriter, r *http.Request, status int, err error) {
	render.Status(r, status)
	render.JSON(w, r, map[string]any{
		"error": err.Error(),
	})
}

func (h *BaseHandler) SendSuccess(w http.ResponseWriter, r *http.Request, status int, message string, results interface{}) {
	render.Status(r, status)
	render.JSON(w, r, map[string]any{
		"message": message,
		"results": results,
	})
}

func ParseRequestBody[T any](r *http.Request, v *util.Validator) (*T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid json body")
	}
	defer r.Body.Close()
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}
	return &data, nil
}


func StructToMap(v any) (map[string]any, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}
	return m, nil
}