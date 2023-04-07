package services

import (
	"yogurt-project/database"
	"yogurt-project/repository"
)

type AuthService struct {
	Username string
	Password string
}

func (a *AuthService) LoginService() (bool, error) {
	authRepo := repository.AuthRepository{
		Db: database.Connect(),
	}
	return authRepo.CheckAuth(a.Username, a.Password)
}

func (a *AuthService) RegisterService() (bool, error) {
	authRepo := repository.AuthRepository{
		Db: database.Connect(),
	}
	return authRepo.CreateNew(a.Username, a.Password)
}
