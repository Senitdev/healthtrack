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

func ParamRoutesPressureMeasure(cx *gin.Engine, db *gorm.DB) {
	r := cx.Group("/api")
	pressureRepo := repository.NewPressureRepository(db)
	pressureService := service.NewPressureService(pressureRepo)
	pressureController := controller.NewPressureMeasureController(pressureService)
	//POST
	r.POST("/pressure", func(ctx *gin.Context) {
		ctx.JSON(200, pressureController.Save(ctx))
		if ctx.Errors != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"Erreur": ""})
		}
	})
	//GET
	r.GET("/pressure/:userId", func(ctx *gin.Context) {
		ids := ctx.Param("userId")
		userId, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou innexistant": userId})
			return
		}
		pressure, err := pressureController.GetAllByUser(userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		ctx.JSON(200, pressure)
	})
	//DELETE
	r.DELETE("/pressure/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ids, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou inexistant": ids})
			return
		}
		if err := pressureController.DeleteById(ids); err != nil {
			ctx.JSON(http.StatusBadRequest, "")
			return
		}
		ctx.JSON(http.StatusNoContent, "")
	})

}
