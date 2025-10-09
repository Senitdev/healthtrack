package repository

import (
	"healthtrack/entity"

	"gorm.io/gorm"
)

type GlucoseRepository interface {
	Save(glucoseMesure entity.GlucoseMeasurement) entity.GlucoseMeasurement
	GetGlucoseAllByUser(userId int) ([]entity.GlucoseMeasurement, error)
	DeleteById(id int) error
}

type glucoseRepository struct {
	BD *gorm.DB
}

// DeleteById implements GlucoseRepository.
func (g *glucoseRepository) DeleteById(id int) error {
	result := g.BD.Delete(&entity.GlucoseMeasurement{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetGlucoseAllByUser implements GlucoseRepository.
func (g *glucoseRepository) GetGlucoseAllByUser(userId int) ([]entity.GlucoseMeasurement, error) {
	var glucoseMeasure []entity.GlucoseMeasurement
	if err := g.BD.Where("user_id = ?", userId).Find(&glucoseMeasure).Error; err != nil {
		return nil, err
	}
	return glucoseMeasure, nil
}

// Save implements GlucoseRepository.
func (g *glucoseRepository) Save(glucoseMesure entity.GlucoseMeasurement) entity.GlucoseMeasurement {
	g.BD.Save(&glucoseMesure)
	return glucoseMesure
}

func NewGlucoseRepository(conne *gorm.DB) GlucoseRepository {
	return &glucoseRepository{
		BD: conne,
	}
}
