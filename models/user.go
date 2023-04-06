package models

import (
	"gorm.io/gorm"
	"time"
	"yogurt-project/database"
	"yogurt-project/pkg/util"
)

type User struct {
	id        int       `gorm:"primary_key;auto_increment" json:"id"`
	username  string    `gorm:"size:255;not null;unique" json:"username"`
	password  string    `gorm:"size:255;not null" json:"password"`
	male      int       `gorm:"size:11;not null" json:"male"`
	role      int       `gorm:"size:11;not null" json:"role"`
	phone     string    `gorm:"size:255;not null" json:"phone"`
	email     string    `gorm:"size:255;not null;unique" json:"email"`
	address   string    `gorm:"size:255;not null" json:"address"`
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// singleton golang
var instance *User

func init() {
	instance = &User{}
}

func construct() *User {
	return instance
}

func (user User) CheckAuth(username, password string) (bool, error) {

	pass, _ := util.HashPassword(password)

	err := database.Db.Where(User{
		username: username,
		password: pass,
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.id > 0 {
		return true, nil
	}

	return false, nil
}

func (user User) CreateNew(username, password string) (bool, error) {
	err := database.Db.Create(User{
		username:  username,
		password:  password,
		male:      0,
		role:      0,
		phone:     "0",
		email:     "not email",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}).First(&user).Error

	if err != nil && err != gorm.ErrInvalidField {
		return false, err
	}

	if user.id > 0 {
		return true, nil
	}

	return false, nil
}
