package validators

import (
	"boilerplate-api/errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// PostValidator structure
type PostValidator struct {
	Validate *validator.Validate
}

// NewPostValidator Register Custom Validators
func NewPostValidator() PostValidator {
	v := validator.New()

	_ = v.RegisterValidation("title", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if value == "" {
			return false
		}
		if len(value) > 255 {
			return false
		}
		return true
	})

	_ = v.RegisterValidation("content", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return value != ""
	})

	return PostValidator{
		Validate: v,
	}
}

func (cv PostValidator) generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	case "phone":
		return fmt.Sprintf("Field '%s' is not valid.", field)
	case "gender":
		return fmt.Sprintf("Field '%s' is not valid.", field)
	case "email":
		return fmt.Sprintf("Field '%s' is not valid.", field)
	case "title":
		return fmt.Sprintf("Field '%s' is required.", field)
	case "content":
		return fmt.Sprintf("Field '%s' is required.", field)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func (cv PostValidator) GenerateValidationResponse(err error) []errors.ErrorContext {
	var validations []errors.ErrorContext
	for _, value := range err.(validator.ValidationErrors) {
		field, rule := value.Field(), value.Tag()
		validation := errors.ErrorContext{Field: field, Message: cv.generateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	return validations
}
