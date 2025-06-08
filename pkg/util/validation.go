package util

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

type Validatable interface {
	Validate() error
}

func init() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func RenderErrors(e validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, err := range e {
		field := err.Field()
		switch err.Tag() {
		case "required":
			errors[field] = field + " is required"
		case "email":
			errors[field] = " invalid email address"
		case "min":
			errors[field] = field + " must be at least " + err.Param() + " characters"
		case "containsany":
			errors[field] = field + " must contain at least on special character " + err.Param()
		default:
			errors[field] = field + " is invalid"
		}
	}
	return errors
}
