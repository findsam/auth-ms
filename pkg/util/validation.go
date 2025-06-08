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

func (v *Validator) ParseValidationErrors(err error) map[string]string {
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"_general": "validation error occurred"}
	}

	errors := make(map[string]string)
	for _, e := range validationErrors {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errors[field] = field + " is required"
		case "email":
			errors[field] = field + " must be a valid email address"
		case "min":
			errors[field] = fmt.Sprintf("%s must be at least %s characters long", field, e.Param())
		case "containsany":
			errors[field] = fmt.Sprintf("%s must contain at least one of: %s", field, e.Param())
		default:
			errors[field] = fmt.Sprintf("%s is invalid", field)
		}
	}
	return errors
}