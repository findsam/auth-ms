package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/findsam/auth-micro/pkg/util"
)

func ParseRequestBody[T any](r *http.Request) (*T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	defer r.Body.Close()
	return &data, nil
}

func ValidateBody[T any](data *T, validator *util.Validator) error {
	if err := validator.Struct(data); err != nil {
		return err
	}
	return nil
}