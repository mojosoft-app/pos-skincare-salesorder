package utils

import (
	"pos-mojosoft-so-service/internal/models"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) []models.ErrorDetail {
	var errors []models.ErrorDetail
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorDetail models.ErrorDetail
			errorDetail.Field = err.Field()

			switch err.Tag() {
			case "required":
				errorDetail.Message = "This field is required"
			case "email":
				errorDetail.Message = "Invalid email format"
			case "min":
				errorDetail.Message = "Minimum length is " + err.Param()
			case "max":
				errorDetail.Message = "Maximum length is " + err.Param()
			case "url":
				errorDetail.Message = "Invalid URL format"
			default:
				errorDetail.Message = "Invalid value"
			}

			errors = append(errors, errorDetail)
		}
	}

	return errors
}
