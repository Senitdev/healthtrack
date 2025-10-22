package handler

import (
	"healthtrack/controller"
	"healthtrack/midllewares"
	"healthtrack/repository"
	"healthtrack/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamRoutesGlucoseMeaure(cx *gin.Engine, db *gorm.DB) {

	glucoseMeasureRepo := repository.NewGlucoseRepository(db)
	glucoseMeasureService := service.NewGlucoseService(glucoseMeasureRepo)
	glucoseMeasureController := controller.NewGlucoseMeasureController(glucoseMeasureService)
	//On protege la route
	r := cx.Group("/api/v1", midllewares.AuthorizeJWT())
	//Save glucose measure
	r.POST("/glucose", func(ctx *gin.Context) {
		ctx.JSON(200, glucoseMeasureController.Save(ctx))
	})
	r.GET("/glucose/:userId", func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		ids, err := strconv.Atoi(userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID invalid"})
		}
		gluceMeasure := glucoseMeasureController.GetGluseMeasureAll(ids)
		ctx.JSON(200, gluceMeasure)

	})
	r.DELETE("/glucose/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "ID manquante"})
			return
		}
		glucoseMeasureController.DeleteById(id)

	})
}
