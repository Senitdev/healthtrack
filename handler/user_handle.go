package handler

import (
	"healthtrack/controller"
	"healthtrack/repository"
	"healthtrack/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamRoutesUser(ct *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	r := ct.Group("/api/v1")
	r.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})
	r.GET("/user/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"ID inconnu ou incorrect": id})
			return
		}
		user, err := userController.GetUserById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Auncun enregistrement pour ce ID": id})
			return
		}
		ctx.JSON(200, user)
	})
	//Get All
	r.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})
	//Get User By Email
	r.GET("/user/email/:email", func(ctx *gin.Context) {
		email := ctx.Param("email")
		if email == "" {
			ctx.JSON(http.StatusBadRequest, "Email manquante")
			return
		}
		user, err := userController.GetUserByEmail(email)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"Error email": err})
			return
		}
		ctx.JSON(200, user)
	})
	//Delete user BY ID
	r.DELETE("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ids, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"ID manquant ou inconnu": ids})
			return
		}
		userController.DeleteById(ids)
		ctx.JSON(http.StatusOK, gin.H{"User supprimer avec succes": ids})

	})

}
