package service

import (
	"context"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/MatheusHenrique129/application-in-go/internal/domain"
	"github.com/MatheusHenrique129/application-in-go/internal/errors"
	"github.com/MatheusHenrique129/application-in-go/internal/repository"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
)

type UserService interface {
	FindByID(ctx context.Context, userID int64) (*domain.User, errors.CustomError)
	Create(ctx context.Context, user domain.CreateUser) (*domain.User, errors.CustomError)
}

type userService struct {
	BaseService

	validateService *ValidateService
	userRepository  repository.UserRepository
}

func (u *userService) FindByID(ctx context.Context, userID int64) (*domain.User, errors.CustomError) {

	user, repoErr := u.userRepository.FindByID(ctx, userID)
	if repoErr != nil {
		return nil, errors.NewRepoErrorResponse(consts.OccurredErrorFindUserMessage, repoErr)
	}

	if user == nil {
		return nil, errors.NewNotFoundCustomError(consts.UserNotFoundMessage)
	}

	return domain.UserToCreateUserDomain(user), nil
}

func (u *userService) Create(ctx context.Context, req domain.CreateUser) (*domain.User, errors.CustomError) {

	if err := u.validateService.validator.Struct(&req); err != nil {
		return nil, errors.NewValidationResponseError(err)
	}

	user := domain.UserToCreateUserModel(&req)

	if repoErr := u.userRepository.Create(ctx, user); repoErr != nil {
		return nil, errors.NewRepoErrorResponse(consts.OccurredErrorCreateUserMessage, repoErr)
	}

	res := domain.UserToCreateUserDomain(user)

	return res, nil
}

func NewUserService(conf *config.Config, validateService *ValidateService, userRepository repository.UserRepository) UserService {
	logger := util.NewLogger("User Service")

	return &userService{
		BaseService:     NewBaseService(conf, logger),
		validateService: validateService,
		userRepository:  userRepository,
	}
}
