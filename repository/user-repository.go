package repository

import (
	"healthtrack/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) entity.User
	FindAll() []entity.User
	GetUserByEmail(email string) (entity.User, error)
	DeleteById(id int) error
	GetUserById(id int) (entity.User, error)
}
type userRepository struct {
	DB *gorm.DB
}

// DeleteById implements UserRepository.
func (u *userRepository) DeleteById(id int) error {
	if resultat := u.DB.Delete(&entity.User{}, id); resultat.Error != nil {
		return resultat.Error
	}
	return nil
}

// FindAll implements UserRepository.
func (u *userRepository) FindAll() []entity.User {
	var user []entity.User
	u.DB.Find(&user)
	return user

}

// GetUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	resultat := u.DB.Where("email", email).Find(&user)
	if resultat.Error != nil {
		return user, resultat.Error
	}
	return user, nil
}

// GetUserById implements UserRepository.
func (u *userRepository) GetUserById(id int) (entity.User, error) {
	var user entity.User
	resultat := u.DB.First(&user, id)
	if resultat.Error != nil {
		return user, resultat.Error
	}
	return user, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(user entity.User) entity.User {
	resultat := u.DB.Save(&user)
	var userVide entity.User
	if resultat.Error != nil {
		return userVide
	}
	return user
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		DB: conn,
	}
}
