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

func ParamRoutesWeight(cx *gin.Engine, db *gorm.DB) {
	weighRepo := repository.NewWeightRepository(db)
	weightService := service.NewWeightService(weighRepo)
	weightController := controller.NewWeightController(weightService)

	r := cx.Group("/api")
	r.POST("/weight", func(ctx *gin.Context) {
		ctx.JSON(200, weightController.Save(ctx))
		if ctx.Errors != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": ctx.Errors})
		}
	})
	//Retourne le poids d'un User
	r.GET("/weight/:userId", func(ctx *gin.Context) {
		idS := ctx.Param("userId")
		id, err := strconv.Atoi(idS)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou inccorect": id})
			return
		}
		weight := weightController.GetWeightByUser(id)
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
