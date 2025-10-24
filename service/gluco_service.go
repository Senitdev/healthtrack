package service

import (
	"healthtrack/entity"
	"healthtrack/repository"
	"time"
)

type GlucoseService interface {
	Save(glucoseMesure entity.GlucoseMeasurement) entity.GlucoseMeasurement
	FindByUser(username string) ([]entity.GlucoseMeasurement, error)
	DeleteById(id int) error
	//Update(glucoseMesure entity.GlucoseMeasurement) (entity.GlucoseMeasurement, error)
}
type glucoseService struct {
	service repository.GlucoseRepository
}

// DeleteById implements GlucoseService.
func (g *glucoseService) DeleteById(id int) error {
	err := g.service.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

// FindByUser implements GlucoseService.
func (g *glucoseService) FindByUser(username string) ([]entity.GlucoseMeasurement, error) {
	var glucoseMeasure []entity.GlucoseMeasurement

	glucoseMeasure, err := g.service.GetGlucoseAllByUser(username)
	if err != nil {
		return glucoseMeasure, err
	}
	return glucoseMeasure, nil
}

// Save implements GlucoseService.
func (g *glucoseService) Save(glucoseMesure entity.GlucoseMeasurement) entity.GlucoseMeasurement {
	glucoseMesure.CreatedAt = time.Now()
	glucoseMesure.Take_at = time.Now()
	glucoseMesure.UpdatedAt = time.Now()
	g.service.Save(glucoseMesure)
	return glucoseMesure
}
func NewGlucoseService(repo repository.GlucoseRepository) GlucoseService {
	return &glucoseService{
		service: repo,
	}
}
