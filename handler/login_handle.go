package handler

import (
	"healthtrack/controller"
	"healthtrack/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamLogin(ctx *gin.Engine, db *gorm.DB) {
	loginService := service.NewLoginService()
	jwtService := service.NewJWTService()
	loginController := controller.NewLoginController(loginService, jwtService)
	r := ctx.Group("/api/v1")
	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	})
}
