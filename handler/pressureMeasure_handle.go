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

func ParamRoutesPressureMeasure(cx *gin.Engine, db *gorm.DB) {
	//On protege la route
	r := cx.Group("/api/v1", midllewares.AuthorizeJWT())
	pressureRepo := repository.NewPressureRepository(db)
	pressureService := service.NewPressureService(pressureRepo)
	pressureController := controller.NewPressureMeasureController(pressureService)
	//POST
	r.POST("/pressure", func(ctx *gin.Context) {
		ctx.JSON(200, pressureController.Save(ctx))
	})
	//GET
	r.GET("/pressure/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		if username == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"username inconnu ou innexistant": username})
			return
		}
		pressure, err := pressureController.GetAllByUser(username)
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
