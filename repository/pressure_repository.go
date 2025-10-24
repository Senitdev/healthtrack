package repository

import (
	"healthtrack/entity"

	"gorm.io/gorm"
)

type PressureRepository interface {
	Save(pressure entity.PressureMeasurement) entity.PressureMeasurement
	GetPressureByUser(username string) ([]entity.PressureMeasurement, error)
	DeletePressure(id int) error
	UpdatePressure(pressure entity.PressureMeasurement, id int) (entity.PressureMeasurement, error)
}
type pressureRepository struct {
	DB *gorm.DB
}

// DeletePressure implements PressureRepository.
func (p *pressureRepository) DeletePressure(id int) error {
	if result := p.DB.Delete(id); result != nil {
		return result.Error
	}
	return nil
}

// GetPressureByUser implements PressureRepository.
func (p *pressureRepository) GetPressureByUser(username string) ([]entity.PressureMeasurement, error) {
	var pressure []entity.PressureMeasurement
	//je recupere le ID de l'user a partir de son username
	var userId int
	if err := p.DB.Table("users").Where("username = ?", username).Select("id").Scan(&userId).Error; err != nil {
		return nil, err
	}
	if err := p.DB.Where("user_id = ?", userId).Find(&pressure).Error; err != nil {
		return nil, err
	}
	return pressure, nil
}

// Save implements PressureRepository.
func (p *pressureRepository) Save(pressure entity.PressureMeasurement) entity.PressureMeasurement {
	var userId int
	if err := p.DB.Table("users").Where("username = ?", pressure.Username).Select("id").Scan(&userId).Error; err != nil {
		return pressure
	}
	pressure.User_id = userId
	p.DB.Save(&pressure)
	return pressure
}

// UpdatePressure implements PressureRepository.
func (p *pressureRepository) UpdatePressure(pressure entity.PressureMeasurement, id int) (entity.PressureMeasurement, error) {
	p.DB.Save(&pressure)
	return pressure, nil
}

func NewPressureRepository(connex *gorm.DB) PressureRepository {
	return &pressureRepository{
		DB: connex,
	}
}
