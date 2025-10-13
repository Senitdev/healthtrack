package server

import (
	"healthtrack/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoute(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	//cors config
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ton front React/Next.js
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))
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
