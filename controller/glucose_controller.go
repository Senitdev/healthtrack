package controller

import (
	"healthtrack/entity"
	"healthtrack/service"

	"github.com/gin-gonic/gin"
)

type GlucoseMeasureController interface {
	Save(ctx *gin.Context) entity.GlucoseMeasurement
	GetGluseMeasureAll(userId int) []entity.GlucoseMeasurement
	Update(ctx *gin.Context, id int) entity.GlucoseMeasurement
	DeleteById(id int) error
}
type glucoseMeasureController struct {
	controller service.GlucoseService
}

// DeleteById implements GlucoseMeasureController.
func (g *glucoseMeasureController) DeleteById(id int) error {
	if err := g.controller.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// GetGluseMeasureAll implements GlucoseMeasureController.
func (g *glucoseMeasureController) GetGluseMeasureAll(userId int) []entity.GlucoseMeasurement {
	var glucoseMeasure []entity.GlucoseMeasurement
	glucoseMeasure, err := g.controller.FindByUser(userId)
	if err != nil {
		return glucoseMeasure
	}
	return glucoseMeasure
}

// Save implements GlucoseMeasureController.
func (g *glucoseMeasureController) Save(ctx *gin.Context) entity.GlucoseMeasurement {
	var glucoseMeasure entity.GlucoseMeasurement
	ctx.BindJSON(&glucoseMeasure)
	g.controller.Save(glucoseMeasure)
	return glucoseMeasure
}

// Update implements GlucoseMeasureController.
func (g *glucoseMeasureController) Update(ctx *gin.Context, id int) entity.GlucoseMeasurement {

	panic("unimplemented")
}

func NewGlucoseMeasureController(service service.GlucoseService) GlucoseMeasureController {
	return &glucoseMeasureController{
		controller: service,
	}
}
