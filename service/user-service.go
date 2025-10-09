package service

import (
	"healthtrack/entity"
	"healthtrack/repository"
)

type UserService interface {
	Save(user entity.User) entity.User
	FindAll() []entity.User
	DeleteUserById(id int) error
	GetUserByEmail(email string) (entity.User, error)
	GetUserById(id int) (entity.User, error)
}
type userService struct {
	service repository.UserRepository
}

// DeleteUserById implements UserService.
func (u *userService) DeleteUserById(id int) error {
	if err := u.service.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements UserService.
func (u *userService) FindAll() []entity.User {
	return u.service.FindAll()
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if user, err := u.service.GetUserByEmail(email); err != nil {
		return user, err
	}
	return user, nil
}

// GetUserById implements UserService.
func (u *userService) GetUserById(id int) (entity.User, error) {
	var user entity.User
	if user, err := u.service.GetUserById(id); err != nil {
		return user, err
	}
	return user, nil
}

// Save implements UserService.
func (u *userService) Save(user entity.User) entity.User {
	u.service.Save(user)
	return user
}

// Contructeur
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		service: repo,
	}
}
