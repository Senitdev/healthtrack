package service

import (
	"healthtrack/entity"
	"healthtrack/repository"
	"time"
)

type WeightService interface {
	Save(weightMeasurement entity.WeightMeasurement) entity.WeightMeasurement
	GetWeightMesureByUser(username string) ([]entity.WeightMeasurement, error)
	DeleteById(id int) error
	UpdateWeightMesurement(weight entity.WeightMeasurement, id int) (entity.WeightMeasurement, error)
}
type weightservice struct {
	service repository.WeightRepository
}

// DeleteById implements WeightService.
func (w *weightservice) DeleteById(id int) error {
	if err := w.service.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// GetWeightMesureByUser implements WeightService.
func (w *weightservice) GetWeightMesureByUser(username string) ([]entity.WeightMeasurement, error) {
	var weightmeasure []entity.WeightMeasurement
	weightmeasure, err := w.service.GetWeightMesureByUser(username)
	if err != nil {
		return weightmeasure, err
	}
	return weightmeasure, nil
}

// Save implements WeightService.
func (w *weightservice) Save(weightMeasurement entity.WeightMeasurement) entity.WeightMeasurement {
	weightMeasurement.TakenAt = time.Now()
	weightMeasurement.CreatedAt = time.Now()
	weightMeasurement.UpdatedAt = time.Now()
	w.service.Save(weightMeasurement)
	return weightMeasurement
}

// UpdateWeightMesurement implements WeightService.
func (w *weightservice) UpdateWeightMesurement(weight entity.WeightMeasurement, id int) (entity.WeightMeasurement, error) {
	panic("unimplemented")
}

func NewWeightService(repo repository.WeightRepository) WeightService {
	return &weightservice{
		service: repo,
	}
}
