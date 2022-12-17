package service

import (
	"regexp"

	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/go-playground/validator/v10"
)

type ValidateService struct {
	validator *validator.Validate
}

func (s *ValidateService) RegisterValidation(name string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return s.validator.RegisterValidation(name, fn, callValidationEvenIfNull...)
}

func (s *ValidateService) Struct(st interface{}) error {
	return s.validator.Struct(st)
}

func (s *ValidateService) validateEmail(fl validator.FieldLevel) bool {
	return isEmailValid(fl.Field().String())
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func NewValidatorService(validator *validator.Validate) *ValidateService {
	vs := &ValidateService{
		validator: validator,
	}

	vs.RegisterValidation(consts.ValidEmailTag, vs.validateEmail)

	return vs
}
