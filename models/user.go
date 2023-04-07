package models

import (
	"time"
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
