package controller

import (
	"healthtrack/entity"
	"healthtrack/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Save(ctx *gin.Context) entity.User
	FindAll() []entity.User
	GetUserByEmail(email string) (entity.User, error)
	DeleteById(id int) error
	GetUserById(id int) (entity.User, error)
}
type userController struct {
	controller service.UserService
}

// DeleteById implements UserController.
func (u *userController) DeleteById(id int) error {
	if err := u.controller.DeleteUserById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements UserController.
func (u *userController) FindAll() []entity.User {
	return u.controller.FindAll()
}

// GetUserByEmail implements UserController.
func (u *userController) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if user, err := u.controller.GetUserByEmail(email); err != nil {
		return user, err
	}
	return user, nil
}

// GetUserById implements UserController.
func (u *userController) GetUserById(id int) (entity.User, error) {
	var user entity.User
	if user, err := u.controller.GetUserById(id); err != nil {
		return user, err
	}
	return user, nil
}

// Save implements UserController.
func (u *userController) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	u.controller.Save(user)
	return user
}

// Contructeur
func NewUserController(repo service.UserService) UserController {
	return &userController{
		controller: repo,
	}
}
