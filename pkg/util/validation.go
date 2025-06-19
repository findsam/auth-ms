package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &Validator{v}
}

func (v *Validator) ParseValidationErrors(err error) []map[string]string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []map[string]string{
			{"key": "_general", "value": "validation error occurred"},
		}
	}

	var errors []map[string]string
	for _, e := range validationErrors {
		field := e.Field()
		var message string

		switch e.Tag() {
		case "required":
			message = field + " is required"
		case "email":
			message = field + " must be a valid email address"
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters long", field, e.Param())
		case "containsany":
			message = fmt.Sprintf("%s must contain at least one of: %s", field, e.Param())
		default:
			message = fmt.Sprintf("%s is invalid", field)
		}

		errors = append(errors, map[string]string{
			"key":   field,
			"value": message,
		})
	}

	return errors
}
