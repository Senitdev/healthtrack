package controller

import (
	"healthtrack/entity"
	"healthtrack/service"

	"github.com/gin-gonic/gin"
)

type WeightMeasureController interface {
	Save(ctx *gin.Context) entity.WeightMeasurement
	GetWeightByUser(userId int) []entity.WeightMeasurement
	UpdateWeightMeasure(ctx *gin.Context, id int) entity.WeightMeasurement
	DeleteWeightMeasureById(id int) error
}
type weightMeasure struct {
	controller service.WeightService
}

// DeleteWeightMeasureById implements WeightMeasureController.
func (w *weightMeasure) DeleteWeightMeasureById(id int) error {
	if err := w.controller.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// GetWeightByUser implements WeightMeasureController.
func (w *weightMeasure) GetWeightByUser(userId int) []entity.WeightMeasurement {
	var weightMeasure []entity.WeightMeasurement
	weightMeasure, err := w.controller.GetWeightMesureByUser(userId)
	if err != nil {
		return weightMeasure
	}
	return weightMeasure
}

// Save implements WeightMeasureController.
func (w *weightMeasure) Save(ctx *gin.Context) entity.WeightMeasurement {
	var weightMeasure entity.WeightMeasurement
	ctx.BindJSON(&weightMeasure)
	w.controller.Save(weightMeasure)
	return weightMeasure
}

// UpdateWeightMeasure implements WeightMeasureController.
func (w *weightMeasure) UpdateWeightMeasure(ctx *gin.Context, id int) entity.WeightMeasurement {
	panic("unimplemented")
}

func NewWeightController(service service.WeightService) WeightMeasureController {
	return &weightMeasure{
		controller: service,
	}
}
