package server

import (
	"healthtrack/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoute(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	//routes gluceMeasure
	handler.ParamRoutesGlucoseMeaure(r, db)
	//Routes Pressure
	handler.ParamRoutesPressureMeasure(r, db)
	//Routes weight
	handler.ParamRoutesWeight(r, db)
	//Routes user
	handler.ParamRoutesUser(r, db)
	return r
}
