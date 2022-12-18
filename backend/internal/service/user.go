package service

import (
	"context"
	"strconv"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/MatheusHenrique129/application-in-go/internal/domain"
	"github.com/MatheusHenrique129/application-in-go/internal/errors"
	"github.com/MatheusHenrique129/application-in-go/internal/repository"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
)

type UserService interface {
	FindByID(ctx context.Context, userID int64) (*domain.User, errors.CustomError)

	Create(ctx context.Context, req domain.CreateUser) (*domain.User, errors.CustomError)
	Update(ctx context.Context, req domain.UpdateUser) (*domain.User, errors.CustomError)
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

func (u *userService) Update(ctx context.Context, req domain.UpdateUser) (*domain.User, errors.CustomError) {
	if err := u.validateService.validator.Struct(&req); err != nil {
		return nil, errors.NewValidationResponseError(err)
	}

	var userID, ParseErr = strconv.ParseInt(req.UserID, consts.DefaultBase, consts.Size64)
	if ParseErr != nil {
		return nil, errors.NewBadRequestBindingResponse(consts.IDCannotStringMessage, ParseErr)
	}

	found, repoErr := u.userRepository.FindByID(ctx, userID)
	if repoErr != nil {
		return nil, errors.NewRepoErrorResponse(consts.OccurredErrorFindUserMessage, repoErr)
	}

	if found == nil {
		return nil, errors.NewNotFoundCustomError(consts.UserNotFoundMessage)
	}

	user := domain.UserToUpdateUserModel(&req)
	user.ID = found.ID

	rowsAffected, err := u.userRepository.Update(ctx, user)
	if err != nil {
		return nil, errors.NewRepoErrorResponse(consts.OccurredErrorUpdateUserMessage, repoErr)
	}

	if rowsAffected == 0 {
		return nil, errors.NewBadRequestCustomError("Could not update any user with those values.")
	}

	return domain.UserToUpdateUserDomain(user), nil
}

func NewUserService(conf *config.Config, validateService *ValidateService, userRepository repository.UserRepository) UserService {
	logger := util.NewLogger("User Service")

	return &userService{
		BaseService:     NewBaseService(conf, logger),
		validateService: validateService,
		userRepository:  userRepository,
	}
}
