package controller

import (
	"healthtrack/dto"
	"healthtrack/service"

	"github.com/gin-gonic/gin"
)

type Loginv2Controller interface {
	Login(ctx *gin.Context) string
}

type loginv2Controller struct {
	loginv2Service service.Loginv2Service
	jWtService     service.JWTService
}

func NewLoginv2Controller(loginService service.Loginv2Service,
	jWtService service.JWTService) Loginv2Controller {
	return &loginv2Controller{
		loginv2Service: loginService,
		jWtService:     jWtService,
	}
}

func (controller *loginv2Controller) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return err.Error()
	}
	result := controller.loginv2Service.GetUserByLogin(credentials.Username, credentials.Password)
	if !result {
		return ""
	}
	return controller.jWtService.GenerateToken(credentials.Username, true)
}
