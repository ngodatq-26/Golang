package services

import "yogurt-project/models"

type AuthService struct {
	Username string
	Password string
}

func (a *AuthService) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *AuthService) Create() (bool, error) {
	return models.CreateNew(a.Username, a.Password)
}
