package repository

import (
	"gorm.io/gorm"
	"time"
	"yogurt-project/models"
	"yogurt-project/pkg/util"
)

type AuthRepository struct {
	Db *gorm.DB
}

func (a *AuthRepository) CreateNew(username, password string) (bool, error) {
	var user models.User
	pass, _ := util.HashPassword(password)
	err := a.Db.Create(&models.User{
		Username:  username,
		Password:  pass,
		Male:      0,
		Role:      0,
		Phone:     "0",
		Email:     "not email",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).First(&user).Error

	if err != nil && err != gorm.ErrInvalidField {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (a *AuthRepository) CheckAuth(username, password string) (bool, error) {
	var user models.User

	err := a.Db.Where(&models.User{
		Username: username,
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return util.DecryptPassword(password, user.Password), nil
	}

	return false, nil
}
