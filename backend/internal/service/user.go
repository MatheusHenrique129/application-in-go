package service

import (
	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/domain"
	"github.com/MatheusHenrique129/application-in-go/internal/errors"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
)

type UserService interface {
	Find(userID domain.URIUser) (*domain.User, errors.CustomError)
}

type userService struct {
	BaseService

	validateService *ValidateService
}

func (u *userService) Find(userID domain.URIUser) (*domain.User, errors.CustomError) {

	if err := u.validateService.validator.Struct(&userID); err != nil {
		return nil, errors.NewValidationResponseError(err)
	}

	user := &domain.User{
		UserID: 21,
		Name:   "Matheus Henrique",
	}

	return user, nil
}

func NewUserService(conf *config.Config, validateService *ValidateService) UserService {
	logger := util.NewLogger("User Service")

	return &userService{
		BaseService:     NewBaseService(conf, logger),
		validateService: validateService,
	}
}
