package repository

import (
	"healthtrack/entity"

	"gorm.io/gorm"
)

type GlucoseRepository interface {
	Save(glucoseMesure entity.GlucoseMeasurement) entity.GlucoseMeasurement
	GetGlucoseAllByUser(username string) ([]entity.GlucoseMeasurement, error)
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
func (g *glucoseRepository) GetGlucoseAllByUser(username string) ([]entity.GlucoseMeasurement, error) {
	var glucoseMeasure []entity.GlucoseMeasurement
	//je recupere le ID de l'user a partir de son username
	var userId int
	if err := g.BD.Table("users").Where("username = ?", username).Select("id").Scan(&userId).Error; err != nil {
		return nil, err
	}
	if err := g.BD.Where("user_id= ?", userId).Find(&glucoseMeasure).Error; err != nil {
		return nil, err
	}
	return glucoseMeasure, nil
}

// Save implements GlucoseRepository.
func (g *glucoseRepository) Save(glucoseMesure entity.GlucoseMeasurement) entity.GlucoseMeasurement {
	var userId int
	if err := g.BD.Table("users").Where("username = ?", glucoseMesure.Username).Select("id").Scan(&userId).Error; err != nil {
		return glucoseMesure
	}
	glucoseMesure.User_id = userId
	g.BD.Save(&glucoseMesure)
	return glucoseMesure
}

func NewGlucoseRepository(conne *gorm.DB) GlucoseRepository {
	return &glucoseRepository{
		BD: conne,
	}
}
