package ginhandler

import (
	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/gin-gonic/gin"
)

type JWTResponse struct {
	Status      string `json:"status"`
	Information string `json:"information"`
}

type AuthHandler struct {
	conf   *config.Config
	logger *util.Logger
}

func (a *AuthHandler) AuthenticateRequest(c *gin.Context) {

}

func NewAuthHandler(conf *config.Config) *AuthHandler {
	return &AuthHandler{
		conf:   conf,
		logger: util.NewLogger("Auth Handler"),
	}
}
