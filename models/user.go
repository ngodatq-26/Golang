package models

import (
	"gorm.io/gorm"
	"time"
	"yogurt-project/database"
	"yogurt-project/pkg/util"
)

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Male      int       `gorm:"size:11;not null" json:"male"`
	Role      int       `gorm:"size:11;not null" json:"role"`
	Phone     string    `gorm:"size:255;not null" json:"phone"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Address   string    `gorm:"size:255;not null" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func CheckAuth(username, password string) (bool, error) {
	var user User
	pass, _ := util.HashPassword(password)

	err := database.ConstructDatabase().Where(&User{
		Username: username,
		Password: pass,
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CreateNew(username, password string) (bool, error) {
	var user User
	err := database.ConstructDatabase().Create(&User{
		Username:  username,
		Password:  password,
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
