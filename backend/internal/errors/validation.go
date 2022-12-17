package errors

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Tag     string `json:"tag" yaml:"tag"`
	Field   string `json:"field" yaml:"field"`
	Message string `json:"message" yaml:"message"`
}

func NewValidationResponseError(err interface{}) CustomError {
	return getValidationError(err)
}

func getValidationError(input interface{}) CustomError {
	switch err := input.(type) {
	case *json.UnmarshalTypeError:
		return validateTypes(err)

	case *strconv.NumError:
		return NewValidationSingleErrorResponse(NewValidationError(consts.BodyTag, consts.TypeTag, consts.RequestFieldsInvalidTypes))

	case validator.ValidationErrors:
		return validateValues(err)

	default:
		return NewBadRequestResponse(consts.InvalidRequestNotParseMessage)
	}
}

func validateTypes(errorType *json.UnmarshalTypeError) CustomError {
	causeList := make([]interface{}, 0)
	field := errorType.Field
	tag := consts.TypeTag

	message := fmt.Sprintf("%s type should be %s", field, errorType.Type.String())

	causeList = append(causeList, getCauseListElement(field, tag, message))

	return NewValidationErrorResponse(causeList)
}

func validateValues(validationErrors validator.ValidationErrors) CustomError {
	causeList := make([]interface{}, 0)

	for _, v := range validationErrors {
		tag := v.Tag()
		field := v.Field()

		message := getErrorMessage(tag)
		causeList = append(causeList, getCauseListElement(field, tag, message))
	}

	return NewValidationErrorResponse(causeList)
}

func getErrorMessage(tag string) string {

	switch tag {

	case consts.ValidEmailTag:
		return consts.EmailInvalidMessage
	default:
		return consts.FieldIsRequiredMessage
	}
}

func getCauseListElement(name, tag, message string) interface{} {
	return ValidationError{
		Field:   name,
		Tag:     tag,
		Message: message,
	}
}

func NewValidationError(field string, tag string, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Tag:     tag,
		Message: message,
	}
}
