package service

import (
	"healthtrack/repository"
)

type Loginv2Service interface {
	GetUserByLogin(username string, password string) bool
}
type loginv2Service struct {
	service repository.LoginRepository
}

// GetUserByLoginAndPassword implements Loginv2Service.
func (l *loginv2Service) GetUserByLogin(username string, password string) bool {
	result := l.service.GetUserByLoginAndPassword(username, password)
	if !result {
		return false
	}
	return true

}

// Login implements Loginv2Service.

func NewLoginV2Service(repo repository.LoginRepository) Loginv2Service {
	return &loginv2Service{
		service: repo,
	}
}
