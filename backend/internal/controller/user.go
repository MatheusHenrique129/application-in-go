package controller

import (
	"net/http"

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

// FindUser List all users.
// @Summary Returns a list of users.
// @Tags products
// @Produce json
// @Success 200 {object} domain.RoleListResource
// @Failure 400 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /v1/user/{user_id} [get]
func (u *UserController) FindUser(c *gin.Context) {
	uriReq := domain.URIUser{}

	if err := c.BindUri(&uriReq); err != nil {
		appErr := errors.NewBadRequestCustomError("Invalid URI values. Please, check that the external_id and namespace parameters are correct.")
		_ = c.Error(appErr)
		return
	}

	res, err := u.userService.Find(uriReq)
	if err != nil {
		u.logger.Errorf(c, "Error find user. userID: %s.", err, uriReq.UserID)
		c.Status(http.StatusInternalServerError)
		c.Error(err)
		return
	}

	u.logger.Info(c, "Successful in getting the products!")
	c.JSON(http.StatusOK, res)
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
		logger:      util.NewLogger("User Controller"),
	}
}
