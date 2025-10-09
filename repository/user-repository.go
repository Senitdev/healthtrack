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
	panic("unimplemented")
}

// FindAll implements UserRepository.
func (u *userRepository) FindAll() []entity.User {
	panic("unimplemented")
}

// GetUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(email string) (entity.User, error) {
	panic("unimplemented")
}

// GetUserById implements UserRepository.
func (u *userRepository) GetUserById(id int) (entity.User, error) {
	panic("unimplemented")
}

// Save implements UserRepository.
func (u *userRepository) Save(user entity.User) entity.User {
	panic("unimplemented")
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		DB: conn,
	}
}
