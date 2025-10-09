package repository

import (
	"healthtrack/entity"

	"gorm.io/gorm"
)

type WeightRepository interface {
	Save(WeightMeasurement entity.WeightMeasurement) entity.WeightMeasurement
	GetWeightMesureByUser(userid int) ([]entity.WeightMeasurement, error)
	DeleteById(id int) error
	UpdateWeightMesurement(weight entity.WeightMeasurement, id int) (entity.WeightMeasurement, error)
}
type weightRepository struct {
	DB *gorm.DB
}

// DeleteById implements WeightRepository.
func (w *weightRepository) DeleteById(id int) error {
	if resultat := w.DB.Delete(id); resultat != nil {
		return resultat.Error
	}
	return nil
}

// GetWeightMesureByUser implements WeightRepository.
func (w *weightRepository) GetWeightMesureByUser(userid int) ([]entity.WeightMeasurement, error) {
	var weightmesure []entity.WeightMeasurement
	if result := w.DB.Where("user_id", userid).Find(&weightmesure).Error; result != nil {
		return weightmesure, result
	}
	return weightmesure, nil
}

// Save implements WeightRepository.
func (w *weightRepository) Save(WeightMeasurement entity.WeightMeasurement) entity.WeightMeasurement {
	w.DB.Save(&WeightMeasurement)
	return WeightMeasurement
}

// UpdateWeightMesurement implements WeightRepository.
func (w *weightRepository) UpdateWeightMesurement(weight entity.WeightMeasurement, id int) (entity.WeightMeasurement, error) {
	panic("unimplemented")
}

func NewWeightRepository(conn *gorm.DB) WeightRepository {
	return &weightRepository{
		DB: conn,
	}
}
