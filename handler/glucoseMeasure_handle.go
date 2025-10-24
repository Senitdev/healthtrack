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
	r.GET("/glucose/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		if username == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Username invalid"})
		}
		glucoseMeasure := glucoseMeasureController.GetGluseMeasureAll(username)
		ctx.JSON(200, glucoseMeasure)

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
