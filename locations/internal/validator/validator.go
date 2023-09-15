package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func getStructTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func New() *validator.Validate {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(getStructTagName)
	return validate
}
