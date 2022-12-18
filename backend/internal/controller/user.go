package controller

import (
	"net/http"
	"strconv"

	"github.com/MatheusHenrique129/application-in-go/internal/consts"
	"github.com/MatheusHenrique129/application-in-go/internal/domain"
	"github.com/MatheusHenrique129/application-in-go/internal/errors"
	"github.com/MatheusHenrique129/application-in-go/internal/service"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	logger      *util.Logger
	userService service.UserService
}

// FindUser Find user by ID.
// @Summary Find user by ID.
// @Tags user
// @Param user_id path integer true "ID of the user"
// @Produce json
// @Success 200 {object} domain.User
// @Failure 400 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /v1/user/{user_id} [get]
func (u *UserController) FindUser(c *gin.Context) {
	uriReq := domain.URIUser{}

	if err := c.BindUri(&uriReq); err != nil {
		appErr := errors.NewBadRequestBindingResponse(consts.InvalidUriValueUserIDMessage, err)
		c.JSON(appErr.Status(), appErr)
		return
	}

	var userID, ParseErr = strconv.ParseInt(uriReq.UserID, consts.DefaultBase, consts.Size64)
	if ParseErr != nil {
		appErr := errors.NewBadRequestBindingResponse(consts.IDCannotStringMessage, ParseErr)
		c.JSON(appErr.Status(), appErr)
		return
	}

	res, err := u.userService.FindByID(c, userID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// Create a new user.
// @Summary Create a new user.
// @Tags user
// @Produce json
// @Param user body domain.CreateUser true "User Information"
// @Success 201 {object} domain.User
// @Failure 400 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /v1/user [post]
func (u *UserController) Create(c *gin.Context) {
	req := domain.CreateUser{}

	if err := c.BindJSON(&req); err != nil {
		appErr := errors.NewBadRequestCustomError(consts.InvalidRequestJsonMessage)
		c.JSON(appErr.Status(), appErr)
		return
	}

	res, err := u.userService.Create(c, req)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
		logger:      util.NewLogger("User Controller"),
	}
}
