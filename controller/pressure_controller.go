package controller

import (
	"healthtrack/entity"
	"healthtrack/service"

	"github.com/gin-gonic/gin"
)

type PressureMeasureController interface {
	Save(ctx *gin.Context) entity.PressureMeasurement
	Update(ctx *gin.Context, id int) (entity.PressureMeasurement, error)
	GetAllByUser(userId int) ([]entity.PressureMeasurement, error)
	DeleteById(id int) error
}
type pressureMeasureController struct {
	controller service.PressureService
}

// DeleteById implements PressureMeasureController.
func (p *pressureMeasureController) DeleteById(id int) error {
	if err := p.controller.DeletePressure(id); err != nil {
		return err

	}
	return nil
}

// GetAllByUser implements PressureMeasureController.
func (p *pressureMeasureController) GetAllByUser(userId int) ([]entity.PressureMeasurement, error) {
	var pressureMeasure []entity.PressureMeasurement
	pressureMeasure, err := p.controller.GetPressureByUser(userId)
	if err != nil {
		return pressureMeasure, err
	}
	return pressureMeasure, nil
}

// Save implements PressureMeasureController.
func (p *pressureMeasureController) Save(ctx *gin.Context) entity.PressureMeasurement {
	var pressureMeasure entity.PressureMeasurement
	ctx.BindJSON(&pressureMeasure)
	p.controller.Save(pressureMeasure)
	return pressureMeasure
}

// Update implements PressureMeasureController.
func (p *pressureMeasureController) Update(ctx *gin.Context, id int) (entity.PressureMeasurement, error) {
	panic("unimplemented")
}

func NewPressureMeasureController(service service.PressureService) PressureMeasureController {
	return &pressureMeasureController{
		controller: service,
	}
}
