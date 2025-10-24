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

func ParamRoutesWeight(cx *gin.Engine, db *gorm.DB) {
	weighRepo := repository.NewWeightRepository(db)
	weightService := service.NewWeightService(weighRepo)
	weightController := controller.NewWeightController(weightService)
	//On protege la route
	r := cx.Group("/api/v1", midllewares.AuthorizeJWT())
	r.POST("/weight", func(ctx *gin.Context) {
		ctx.JSON(200, weightController.Save(ctx))
		if ctx.Errors != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": ctx.Errors})
		}
	})
	//Retourne le poids d'un User
	r.GET("/weight/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		if username == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"username inconnu ou incorrect": username})
			return
		}
		weight := weightController.GetWeightByUser(username)
		ctx.JSON(200, weight)
	})
	//Delete
	r.DELETE("/weight/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou incorrect": id})
		}
		weightController.DeleteWeightMeasureById(id)
		ctx.JSON(http.StatusNoContent, "")
	})

}
