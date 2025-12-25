package main

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func noBlankValidator(flv validator.FieldLevel) bool {
	value := flv.Field().String()
	return strings.TrimSpace(value) != ""
}
