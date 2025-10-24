package service

import (
	"healthtrack/entity"
	"healthtrack/repository"
	"time"
)

type PressureService interface {
	Save(pressure entity.PressureMeasurement) entity.PressureMeasurement
	GetPressureByUser(username string) ([]entity.PressureMeasurement, error)
	DeletePressure(id int) error
	UpdatePressure(pressure entity.PressureMeasurement, id int) (entity.PressureMeasurement, error)
}
type pressureService struct {
	service repository.PressureRepository
}

// DeletePressure implements PressureService.
func (p *pressureService) DeletePressure(id int) error {
	if err := p.service.DeletePressure(id); err != nil {
		return err
	}
	return nil
}

// GetPressureByUser implements PressureService.
func (p *pressureService) GetPressureByUser(username string) ([]entity.PressureMeasurement, error) {
	var pressure []entity.PressureMeasurement
	pressure, err := p.service.GetPressureByUser(username)
	if err != nil {
		return pressure, err
	}
	return pressure, nil
}

// Save implements PressureService.
func (p *pressureService) Save(pressure entity.PressureMeasurement) entity.PressureMeasurement {
	pressure.Created_At = time.Now()
	pressure.Taken_at = time.Now()
	pressure.Updated_At = time.Now()
	p.service.Save(pressure)
	return pressure
}

// UpdatePressure implements PressureService.
func (p *pressureService) UpdatePressure(pressure entity.PressureMeasurement, id int) (entity.PressureMeasurement, error) {
	panic("unimplemented")
}

func NewPressureService(repo repository.PressureRepository) PressureService {
	return &pressureService{
		service: repo,
	}
}
