package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
)

func ParseError(err error) []string {

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			switch e.Tag() {
			case "required":
				errorMessages[i] = fmt.Sprintf("The field %s is required", e.StructField())
			}
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("The field %s must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}
