package handler

import (
	"healthtrack/controller"
	"healthtrack/repository"
	"healthtrack/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamLogin2(ctx *gin.Engine, db *gorm.DB) {
	login2repo := repository.NewLoginRepository(db)
	login2service := service.NewLoginV2Service(login2repo)
	jwtService := service.NewJWTService()
	login2controller := controller.NewLoginv2Controller(login2service, jwtService)

	r := ctx.Group("/api/v1")
	r.POST("/auth", func(ctx *gin.Context) {
		token := login2controller.Login(ctx)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	})
}
